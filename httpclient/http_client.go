package httpclient

import (
	"bytes"
	"errors"
	"github.com/ztaoing/infra/lb"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultHttpTimeout = 30 * time.Second
)

var parseUrl = url.Parse

type Option struct {
	Timeout time.Duration //超时时间
}
type HttpClient struct {
	client *http.Client
	Option Option
	apps   *lb.Apps
}

func NewHttpClient(apps *lb.Apps, opt *Option) *HttpClient {
	c := &HttpClient{
		apps: apps,
	}
	if opt == nil {
		c.Option = Option{
			Timeout: defaultHttpTimeout,
		}
	} else {
		c.Option = *opt
	}
	//接下来初识化client
	c.client = &http.Client{
		Timeout: c.Option.Timeout,
	}
	return c
}

func (h *HttpClient) NewRequest(method, url string, body io.Reader, headers http.Header) (*http.Request, error) {
	if method == "" {
		method = http.MethodGet
	}
	//解析url
	u, err := parseUrl(url)
	if err != nil {
		return nil, err
	}

	//从解析后的URL中提取微服务名称
	name := u.Host
	//通过微服务名称从本地服务注册表中查询应用和应用实例列表
	app := h.apps.Get(name)
	if app == nil {
		return nil, errors.New("未查询到可用的微服务应用:" + name + "," + url)
	}

	//通过负载均衡算法从应用实例列表中选择一个实例
	ins := app.Get(url)
	if ins == nil {
		return nil, errors.New("未查询到可用的实例:" + name + "," + url)
	}
	//使用选择的实例的ip和port替换URL中的域名部分
	u.Host = ins.Address
	url = u.String()
	//使用新构造的URL创建一个request请求
	r, err := http.NewRequest(method, url, body)
	//将headers添加到request中
	if len(headers) > 0 {
		for key, value := range headers {
			for _, val := range value {
				r.Header.Add(key, val)
			}
		}
	}
	return r, err
}

//通用的http的请求方法
func (h *HttpClient) Do(r *http.Request) (*http.Response, error) {
	resp, err := h.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//重新封装body
	//正常情况下body关闭后就不能再次读取了，但通过NopCloser处理的body关闭后可以再次读取
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return resp, err
}
