/*
Created on 2018/5/22 13:55

author: ChenJinLong

Content:

观察者模式是使用频率最高的设计模式之一，它用于建立一种对象与对象之间的依赖关系，
一个对象发生改变时将自动通知其他对象，其他对象将相应作出反应。
在观察者模式中，发生改变的对象称为观察目标，
而被通知的对象称为观察者，一个观察目标可以对应多个观察者，而且这些观察者之间可以没有任何相互联系，
可以根据需要增加和删除观察者，使得系统更易于扩展。
*/
package observer

import "fmt"

//抽象观察类
type Observer interface {
	getName() string
	setName(name string)
	help()  //声明支援盟友的方法
	beAttacked(acc AllyControlCenter) //声明遭受攻击的方法
}

//战队成员类：具体观察者类
type Player struct {
	name string
}

func (p *Player) getName() string{
	return p.name
}

func (p *Player) setName(name string) {
	p.name = name
}

func (p *Player)help(){
	fmt.Println("坚持住",p.name,"来救你!")
}
func (p *Player)beAttacked(acc AllyControlCenter) {
	fmt.Println(p.name,"被攻击!")
	acc.notifyObserver(p.name)

}

//战队控制中心类，抽象目标类
type AllyControlCenter interface {
	setAllyName(allyName string)
	getAllyName() string
	join(obs Observer) //注册观察对象
	quit(obs Observer) // 注销观察对象
	notifyObserver(name string)

}


//具体战队控制中心,具体目标类
type ConcreteAllyControlCenter struct {
	observers map[string]Observer
	allyName string
}

func NewConcreteAllyControlCenter(allayName string) *ConcreteAllyControlCenter{
	fmt.Println(allayName,"战队组建成功！")
	fmt.Println("----------------------------")
	c := new(ConcreteAllyControlCenter)
	c.observers = map[string]Observer{}
	c.setAllyName(allayName)
	return c
}

func (c *ConcreteAllyControlCenter) setAllyName(allyName string) {
	c.allyName = allyName
}

func (c *ConcreteAllyControlCenter)getAllyName() string{
	return c.allyName
}

func (c *ConcreteAllyControlCenter) join(obs Observer) {
	fmt.Println(obs.getName(),"加入",c.allyName,"战队！")
	c.observers[obs.getName()] = obs
}

func (c *ConcreteAllyControlCenter)quit(obs Observer) {
	fmt.Println(obs.getName(),"退出",c.allyName,"战队")
	delete(c.observers,obs.getName())
}

func (c *ConcreteAllyControlCenter)notifyObserver(name string) {
	fmt.Println(c.allyName,"战队紧急通知,盟友",name,"遭受敌人攻击！")
	for k,v := range c.observers {
		if k != name {
			v.help()
		}
	}
}











