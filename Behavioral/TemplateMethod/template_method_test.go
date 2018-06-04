package template_method

import (
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	handle := new(HandleInterest)
	handle.IAccount = new(SavingAccount)
	handle.Handle("张无忌", "123456")


}