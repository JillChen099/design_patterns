package State

import (
	"testing"
)

func TestNewScreen(t *testing.T) {
	newScreen := NewScreen()
	newScreen.OnClick()
	newScreen.OnClick()
	newScreen.OnClick()
}
