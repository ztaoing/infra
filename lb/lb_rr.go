package lb

import (
	"math/rand"
	"sync/atomic"
)

//简单轮询算法
var _ Balancer = new(RoundRobinBalancer)

type RoundRobinBalancer struct {
	ct uint32 //计数器
}

// 计数器+1自增 然后取模
func (r *RoundRobinBalancer) Next(key string, hosts []*ServerInstance) *ServerInstance {
	if len(hosts) == 0 {
		return nil
	}
	//自增
	counter := atomic.AddUint32(&r.ct, 1)
	//根据host的长度来取模
	index := int(counter) % len(hosts)
	//通过索引来获得实例
	instance := hosts[index]
	return instance
}

//随机轮询算法
var _ Balancer = new(RandomBalancer)

type RandomBalancer struct {
}

//随机数字
func (r *RandomBalancer) Next(key string, hosts []*ServerInstance) *ServerInstance {
	if len(hosts) == 0 {
		return nil
	}
	//返回非负的int32
	counter := rand.Int31()
	index := int(counter) % len(hosts)
	instance := hosts[index]
	return instance
}
