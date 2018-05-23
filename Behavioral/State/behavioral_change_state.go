/*
Created on 2018/5/18 17:52

author: ChenJinLong

Content: 由环境类改变状态
*/
package State

import (
	"fmt"
	"reflect"
	"strings"
)

type IScreenState interface {
	display()
}

type NormalScreenState struct {

}
func (n *NormalScreenState)display() {

	fmt.Println("正常大小！")
}

type LargerScreenState struct {

}
func (l *LargerScreenState) display() {
	fmt.Println("二倍大小！")
}

type LargestScreenState struct {

}
func (l *LargestScreenState) display() {
	fmt.Println("四倍大小")
}


type Screen struct {
	normalState IScreenState
	largerState IScreenState
	largestState IScreenState
	currentState IScreenState
}

func NewScreen() *Screen {
	s := new(Screen)
	s.normalState = new(NormalScreenState)
	s.largerState = new(LargerScreenState)
	s.largestState = new(LargestScreenState)
	s.currentState = s.normalState
	s.currentState.display()
	return s
}


func (s *Screen) setState(state IScreenState) {

	s.currentState = state
}

func (s *Screen) OnClick() {
	stateKind := reflect.TypeOf(s.currentState).String()
	stateKindName := strings.Split(stateKind, ".")[1]
	switch stateKindName {
	case "NormalScreenState":
		s.setState(s.largerState)
		s.currentState.display()
	case "LargerScreenState":
		s.setState(s.largestState)
		s.currentState.display()
	case "LargestScreenState":
		s.setState(s.normalState)
		s.currentState.display()
	}

}

