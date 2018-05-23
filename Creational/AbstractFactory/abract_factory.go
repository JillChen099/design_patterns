/*
Created on 2018/5/17 16:42

author: ChenJinLong

Content:
抽象工厂模式为创建一组对象提供了一种解决方案。
与工厂方法模式相比，抽象工厂模式中的具体工厂不只是创建一种产品，它负责创建一族产品。
*/
package AbstractFactory

import (
	"fmt"
)


//产品A接口
type Button interface {
	display()
}

type SpringButton struct {}

func (s *SpringButton)display() {
	fmt.Println("显示浅绿色按钮！")
}
type SummerButton struct {}

func (s *SummerButton) display() {
	fmt.Println("显示浅蓝色按钮！")
}

//产品B接口
type TextFiled interface {
	display()
}

type SpringTextField struct {

}

func (s *SpringTextField) display() {
	fmt.Println("显示浅绿色文本框！")
}

type SummerTextFiled struct {

}

func (s *SummerTextFiled) display() {
	fmt.Println("显示浅蓝色文本框！")
}

//产品C接口
type ComboBox interface {
	display()
}

type SpringComboBox struct {

}

func (s *SpringComboBox) display() {
	fmt.Println("显示浅绿色组合框！")
}

type SummerComboBox struct {

}
func (s *SummerComboBox) display() {
	fmt.Println("显示浅蓝色组合框！")
}



//抽象工厂接口
type SkinFactory interface {
	createButton() Button
	createTextFiled()  TextFiled
	createComboBox()  ComboBox
}


//具体工厂A
type SpringSkinFactory struct {

}

func (s *SpringSkinFactory) createButton() Button {
	b := new(SpringButton)
	return b
}

func (s *SpringSkinFactory) createTextFiled() TextFiled {
	t := new(SpringTextField)
	return t
}

func (s *SpringSkinFactory) createComboBox() ComboBox {
	c := new(SpringComboBox)
	return c
}


//具体工厂B

type SummerSkinFactory struct {

}

func (s *SummerSkinFactory) createButton() Button{
	b := new(SummerButton)
	return b
}

func (s *SummerSkinFactory) createTextFiled() TextFiled {
	t := new(SummerTextFiled)
	return t
}

func (s *SummerSkinFactory) createComboBox() ComboBox {
	c := new(SummerComboBox)
	return c
}






