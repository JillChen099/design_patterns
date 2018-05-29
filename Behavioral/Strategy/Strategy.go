/*
Created on 2018/5/29 9:45

author: ChenJinLong

Content:

策略模式(Strategy Pattern)：定义一系列算法类，将每一个算法封装起来，并让它们可以相互替换，
策略模式让算法独立于使用它的客户而变化，也称为政策模式(Policy)。策略模式是一种对象行为型模式。
*/
package strategy

import "fmt"

type Discount interface {
	calculate(price float64) float64
}

type StudentDiscount struct {

}

func (self *StudentDiscount)calculate(price float64) float64 {
	fmt.Println("学生票：")
	return price * 0.8
}

type ChildrenDiscount struct {
}

func (self *ChildrenDiscount) calculate(price float64) float64 {
	fmt.Println("儿童票：")
	return price - 10

}

type VIPDiscount struct {

}
func (self *VIPDiscount)calculate(price float64) float64 {
	fmt.Println("VIP票：")
	fmt.Println("增加积分！")
	return price * 0.5
}


type MovieTicker struct {
	price float64
	discount Discount
}

func (m *MovieTicker) setPrice(price float64) {
	m.price = price
}

func (m *MovieTicker) setDiscount(d Discount) {
	m.discount = d

}

func (m *MovieTicker)getPrice() float64 {
	return m.discount.calculate(m.price)
}



