/*
Created on 2018/5/21 13:45

author: ChenJinLong

Content:
确保某一个类只有一个实例，而且自行实例化并向整个系统提供这个实例，
这个类称为单例类，它提供全局访问的方法。单例模式是一种对象创建型模式。
*/
package Singleton

import (
	"sync"
	"math/rand"
)
type List struct {
	mapList map[int]string
}

func (l *List) add(s string) {
	listLen := len(l.mapList)
	l.mapList[listLen] = s
}

func (l *List) remove(s string) {
	for k,v := range l.mapList {
		if v == s {
			delete(l.mapList,k)
		}
	}
}

func (l *List) get(i int)string {
	return l.mapList[i]

}

var instance *LoadBalancer
var once sync.Once

type LoadBalancer struct {
	serverList  *List
}

func GetLoadBalancer() *LoadBalancer{

	once.Do(func() {
		instance = &LoadBalancer{&List{make(map[int]string)}}
	})
	return instance
}


func (l *LoadBalancer)addServer( server string) {
	l.serverList.add(server)
}

func (l *LoadBalancer) removeServer( server string) {
	l.serverList.remove(server)
}

func (l *LoadBalancer)getServer() string {
	randomInt := rand.Intn(len(l.serverList.mapList))
	res := l.serverList.get(randomInt)
	return res
}




