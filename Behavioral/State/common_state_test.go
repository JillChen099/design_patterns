package State

import (
	"testing"
)

func TestNewSwitch(t *testing.T) {
	s1 := NewSwitch("开关1")
	s2 := NewSwitch("开关2")
	s1.On()
	s2.On()
	s1.Off()
	s2.Off()
	s2.On()
	s1.On()


}
