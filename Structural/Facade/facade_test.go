package facade

import (
	"testing"
)

func TestFacade(t *testing.T) {
	facadeSample := new(EncryptFacade)
	facadeSample.FileEncrypt("A","123")

}
