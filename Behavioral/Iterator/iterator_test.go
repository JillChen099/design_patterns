package iterator

import (
	"testing"
	"fmt"
)

func TestIterator(t *testing.T) {

	products := NewProductList()
	products.addObject("倚天剑")
	products.addObject("屠龙刀")
	products.addObject("断肠草")
	products.addObject("葵花宝典")
	products.addObject("四十二章经")

	iterator := products.createIterator()

	fmt.Println("正向遍历：")
	for !iterator.isLast() {
		fmt.Print(iterator.getNextItem(),",")
		iterator.next()
	}
	fmt.Println("")
	fmt.Println("----------------------")
	fmt.Println("反向遍历")

	for !iterator.isFirst() {
		fmt.Print(iterator.getPreviousItem(),",")
		iterator.previous()
	}

}
