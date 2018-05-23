package adapter

import (
	"testing"
	"fmt"
	"github.com/modern-go/reflect2"
)

func TestAdapter(t *testing.T) {
	var operation ScoreOperation
	operation = new(OperationAdapter)
	scores := []int{84,76,50,69,90,91,88,96}
	fmt.Println("成绩排序结果：")
	operation.sort(scores)
	fmt.Println(scores)
	fmt.Println("查找成绩90：")
	isExistence := operation.search(scores,90)
	if isExistence == 1 {
		fmt.Println("找到成绩90")
	}else {
		fmt.Println("没有找到成绩")
	}

	ty := reflect2.TypeByName("adapter.OperationAdapter").New().(ScoreOperation)
	fmt.Println("查找成绩68：")
	isExistence = ty.search(scores,68)
	if isExistence == 1 {
		fmt.Println("找到成绩68")
	}else {
		fmt.Println("没有找到成绩")
	}




}