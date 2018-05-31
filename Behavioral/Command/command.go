/*
Created on 2018/5/31 10:20

author: ChenJinLong

Content:
命令模式的核心在于引入了命令类，
通过命令类来降低发送者和接收者的耦合度，
请求发送者只需指定一个命令对象，
再通过命令对象来调用请求接收者的处理方法.
*/
package command

import (
	"fmt"
)
//功能键设置窗口类
type FBSettingWindow struct {
	title string
	fbList []*FunctionButton
}

func (f *FBSettingWindow) setTitle(title string) {
	f.title = title
}

func (f *FBSettingWindow) getTitle() string {
	return f.title
}

func (f *FBSettingWindow) addFunctionButton(fb *FunctionButton) {
	f.fbList = append(f.fbList,fb)

}

func (f *FBSettingWindow) removeFunctionButton(fb *FunctionButton) {
	for k,v := range f.fbList {
		if k == 0 && v == fb {
			f.fbList = f.fbList[1:]
			break
		}
		if k == len(f.fbList)-1 && v == fb{
			f.fbList = f.fbList[:len(f.fbList)-2]
			break
		}
		if v == fb {
			f.fbList = append(f.fbList[:k-1],f.fbList[k+1:]...)
			break
		}
	}
}


func (f *FBSettingWindow) display() {
	fmt.Println("显示窗口：",f.title)
	fmt.Println("显示功能键：")
	for _,v := range f.fbList {
		fmt.Println(v.getName())
	}
	fmt.Println("-----------------")

}

//---------------------------------------------//
//功能键类：请求发送者
type FunctionButton struct {
	name string
	command Command
}

func (f *FunctionButton)setName(name string) {
	f.name = name
}

func (f *FunctionButton)getName() string {
	return f.name
}


func (f *FunctionButton) setCommand(command Command) {
	f.command = command
}

func (f *FunctionButton) onClick() {
	fmt.Println("点击功能键：")
	f.command.execute()
}
//-----------------------------------------------------//

//抽象命令类
type Command interface {
	execute()
}

//帮助命令类：具体命令类
type HelpCommand struct {
	hhObj  *HelpHandler
}


func (h *HelpCommand) execute() {
	h.hhObj.display()
}

type MinimizeCommand struct {
	whObj *WindowHanlder
}


func (m *MinimizeCommand) execute(){
	m.whObj.display()
}

//---------------------------------//
//处理类：请求接收者
type HelpHandler struct {
}

func (h *HelpHandler) display() {
	fmt.Println("显示帮助文档！")
}

type WindowHanlder struct {

}

func (w *WindowHanlder) display() {
	fmt.Println("将窗口最小化至托盘！")

}





