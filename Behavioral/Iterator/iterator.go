/*
Created on 2018/5/22 10:56

author: ChenJinLong

Content:
迭代器模式提供一种方法来访问聚合对象，而不用暴露这个对象的内部表示。迭代器模式是一种对象行为型模式
*/
package iterator

//抽象聚合类
type AbstractObjectList interface {
	addObject(obj string)
	removeObject(obj string)
	getObjects() map[int]string
	createIterator() AbstractIterator
}

func NewProductList() *ProductList {
	p := new(ProductList)
	p.Index = 0
	p.Objects = map[int]string{}
	return p
}

//具体聚合类
type ProductList struct {
	Objects map[int]string
	Index   int
}

func (p *ProductList) addObject(obj string) {
	p.Objects[p.Index] = obj
	p.Index ++
}

func (p *ProductList)removeObject(obj string) {
	for k,v := range p.Objects {
		if v == obj {
			delete(p.Objects,k)
		}
	}
}

func (p *ProductList) getObjects() map[int]string {
	return p.Objects

}


func (p *ProductList)createIterator() AbstractIterator {
	productIterator := new(ProductIterator)
	productIterator.products = p.Objects
	productIterator.cursor1 = 0
	productIterator.cursor2 = len(p.Objects) - 1
	return productIterator
}


//抽象迭代器类
type AbstractIterator interface {
	next()
	isLast()  bool
	previous()
	isFirst() bool
	getNextItem() string
	getPreviousItem() string

}


//具体迭代器
type ProductIterator struct {
	products map[int]string
	cursor1 int
	cursor2 int
}

func (p *ProductIterator) next() {
	if p.cursor1 < len(p.products) {
		p.cursor1 ++
	}
}

func (p *ProductIterator) isLast() bool {
	return (p.cursor1 == len(p.products))

}

func (p *ProductIterator)previous() {
	if p.cursor2 > -1 {
		p.cursor2 --
	}
}

func (p *ProductIterator)isFirst() bool {
	return (p.cursor2 == -1)

}

func (p *ProductIterator)getNextItem() string {
	return p.products[p.cursor1]
}

func (p *ProductIterator)getPreviousItem() string {
	return p.products[p.cursor2]
}








