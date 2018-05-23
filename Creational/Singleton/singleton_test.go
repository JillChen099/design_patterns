package Singleton

import (
	"testing"
	"fmt"
)

func TestSingleton(t *testing.T) {
	balancer1 := GetLoadBalancer()
	balancer2 := GetLoadBalancer()
	balancer3 := GetLoadBalancer()
	if (balancer1 == balancer2 && balancer1 == balancer3 && balancer2 == balancer3) {
		fmt.Println("服务器负载均衡器具有唯一性")
	}
	balancer1.addServer("Server 1")
	balancer1.addServer("Server 2")
	balancer1.addServer("Server 3")
	balancer1.addServer("Server 4")
	balancer2.addServer("Server 5")
	balancer3.removeServer("Server 5")
	for i:=0;i<10 ;i++  {
		server := balancer1.getServer()
		fmt.Printf("请求分发至%s服务器\n",server)
	}

}
