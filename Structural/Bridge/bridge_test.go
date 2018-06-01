package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {

	image := new(JPGImage)
	imp := new(WindowsImp)
	image.setImageTmp(imp)
	image.parseFile("小龙女")


}