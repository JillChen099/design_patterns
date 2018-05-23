package AbstractFactory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	factory := new(SpringSkinFactory)
	button := factory.createButton()
	textField := factory.createTextFiled()
	comboBox := factory.createComboBox()
	button.display()
	textField.display()
	comboBox.display()

}

