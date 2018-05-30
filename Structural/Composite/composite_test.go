package composite

import (
	"testing"
	"fmt"
)

func TestComposite(t *testing.T) {
	folder1 := new(Floder)
	folder1.name = "Sunny的资料"

	folder2 := new(Floder)
	folder2.name = "图像文件夹"

	folder3 := new(Floder)
	folder3.name = "文本文件夹"

	file1 := new(ImageFile)
	file1.name = "小龙女.jpg"

	file2 := new(ImageFile)
	file2.name = "张无忌.gif"

	file3 := new(TextFile)
	file3.name = "九阴真经.txt"

	file4 := new(TextFile)
	file4.name = "葵花宝典.txt"

	folder1.add(file4)
	folder2.add(file1)
	folder2.add(file2)
	folder3.add(file3)
	folder1.add(folder2)
	folder1.add(folder3)

	folder1.killVirus()
	fmt.Println("-------------------------")
	folder2.killVirus()

}