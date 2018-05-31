package command

import (
	"testing"
)

func TestCommand(t *testing.T) {

	fbsw := new(FBSettingWindow)
	fbsw.title = "功能键设置"


	fb1 := new(FunctionButton)
	fb1.name = "功能键1"

	command1 := new(HelpCommand)
	fb1.setCommand(command1)



	fb2 := new(FunctionButton)
	fb2.name = "功能键2"

	command2 := new(MinimizeCommand)
	fb2.setCommand(command2)


	fbsw.addFunctionButton(fb1)
	fbsw.addFunctionButton(fb2)
	fbsw.display()

	fb1.onClick()
	fb2.onClick()



}