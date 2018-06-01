/*
Created on 2018/6/1 9:36

author: ChenJinLong

Content:
用于处理多维度变化的设计模式——桥接模式。
桥接模式是一种很实用的结构型设计模式，如果软件系统中某个类存在两个独立变化的维度，
通过该模式可以将这两个维度分离出来，使两者可以独立扩展，让系统更加符合“单一职责原则”。
与多层继承方案不同，它将两个独立变化的维度设计为两个独立的继承等级结构，
并且在抽象层建立一个抽象关联，该关联关系类似一条连接两个独立继承结构的桥，故名桥接模式。
*/
package bridge

import "fmt"

type Matrix struct {

}
//-----------------------------//
type ImageImp interface {
	doPaint(m *Matrix)
}

type WindowsImp struct {

}

func (w *WindowsImp) doPaint(m *Matrix) {
	fmt.Println("在windwo系统中显示图像！")
}

type LinuxImp struct {

}

func (l *LinuxImp) doPaint(m *Matrix) {
	fmt.Println("在Linux系统中显示图像！")
}

type UnixImp struct {

}

func (u *UnixImp)doPaint(m *Matrix) {
	fmt.Println("在Unix系统中显示图像！")
}

//-----------------------------------------//


type Image struct {
	imp ImageImp
}

func (i *Image) setImageTmp(imp ImageImp) {
	i.imp = imp
}

func (i *Image) parseFile(filename string) {

}


type JPGImage struct {
	Image
}

func (j *JPGImage) parseFile(filename string) {
	m := new(Matrix)
	j.imp.doPaint(m)
	fmt.Println(filename,",格式为JPG。")
}

type PNGImage struct {
	Image
}

func (j *PNGImage) parseFile(filename string) {
	m := new(Matrix)
	j.imp.doPaint(m)
	fmt.Println(filename,",格式为PNG。")
}


type GIFImage struct {
	Image
}

func (j *GIFImage) parseFile(filename string) {
	m := new(Matrix)
	j.imp.doPaint(m)
	fmt.Println(filename,",格式为GIF。")
}


