package lb

import "hash/crc32"

//hash算法

var _ Balancer = new(HashBalancer)

type HashBalancer struct {
}

func (h *HashBalancer) Next(key string, hosts []*ServerInstance) *ServerInstance {
	//key 可以为ip、域名、path等
	if len(hosts) == 0 {
		return nil
	}
	//计算出hash值
	counter := crc32.ChecksumIEEE([]byte(key))

}
