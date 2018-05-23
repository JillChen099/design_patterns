/*
Created on 2018/5/18 16:26

author: ChenJinLong

Content: 
*/
package State

import "fmt"

type Switch struct {
	state      *State
	name string
}


var S = State{1}


func NewSwitch(name string) *Switch {
	s := new(Switch)
	s.name = name
	s.state = &S
	return s
}



func (s *Switch) On() {
	fmt.Print(s.name)
	s.state.on()
}

func (s *Switch) Off() {
	fmt.Print(s.name)
	s.state.off()
}


type State struct {
	flag int

}

func (s *State) on() {
	if s.flag == 1 {
		fmt.Print("已经打开\n")
	}else {
		s.flag =1
		fmt.Print("打开\n!")
	}

}

func (s *State)off() {

	if s.flag == 1 {
		s.flag = 2
		fmt.Print("关闭\n")
	}else {
		fmt.Print("已经关闭\n!")
	}

}






