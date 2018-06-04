/*
Created on 2018/6/4 9:23

author: ChenJinLong

Content:
定义：定义一个操作中算法的框架，而将一些步骤延迟到子类中。模板方法模式使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。
在模板方法模式中，将实现功能的每一个步骤所对应的方法称为基本方法（例如“点单”、“吃东西”和“买单”），
而调用这些基本方法同时定义基本方法的执行次序的方法称为模板方法（例如“请客”）。
*/
package template_method

import "fmt"


type IAccount interface {
	Validate(account ,password string) bool
	CalculateInterest()
	Display()

}

type HandleInterest struct {
	IAccount
}

func (h *HandleInterest) Handle(account,password string) {
	if !h.Validate(account,password) {
		fmt.Println("账号或密码错误！")
	}else {
		h.CalculateInterest()
		h.Display()
	}

}


//-----------------------------------//

type Account struct {

}

func (a *Account) Validate(account ,password string) bool{
	fmt.Println("账号：",account)
	fmt.Println("密码：",password)

	if account == "张无忌" && password == "123456" {
		return true

	}else {
		return false
	}
}


func (a *Account)Display() {
	fmt.Println("显示利息！")
}


//----------------------------------------------------//
//活期账户
type CurrentAccount struct {
	Account
}

func (c *CurrentAccount) CalculateInterest() {
	fmt.Println("按活期利率计算利息！")
}

//定期账户
type SavingAccount struct {
	Account
}
func (s *SavingAccount) CalculateInterest() {
	fmt.Println("按定期利率计算利息！")
}










