/*
Created on 2018/5/30 9:34

author: ChenJinLong

Content:

组合多个对象形成树形结构以表示具有“整体—部分”关系的层次结构。
组合模式对单个对象（即叶子对象）和组合对象（即容器对象）的使用具有一致性，
组合模式又可以称为“整体—部分”(Part-Whole)模式，它是一种对象结构型模式
*/
package composite

import "fmt"

type AbstractFile interface {
	add(file AbstractFile)
	remove(file AbstractFile)
	getChild(i int) AbstractFile
	killVirus()
}

//------------------------------------------------//

type ImageFile struct {
	name string
}

func (self *ImageFile) ImageFile(name string) {
	self.name = name
}

func (self *ImageFile) add(file AbstractFile) {
	fmt.Println("对不起，不支持改方法！")
}

func (self *ImageFile) remove(file AbstractFile) {
	fmt.Println("对不起，不支持改方法！")
}

func (self *ImageFile)getChild(i int) AbstractFile {
	fmt.Println("对不起，不支持该方法！")
	return nil
}

func (self *ImageFile) killVirus(){
	fmt.Println("对图像文件",self.name,"进行杀毒")
}

//____________________________________________________//
type TextFile struct {
	name string
}

func (self *TextFile) TextFile(name string) {
	self.name = name
}

func (self *TextFile) add(file AbstractFile) {
	fmt.Println("对不起，不支持改方法！")
}

func (self *TextFile) remove(file AbstractFile) {
	fmt.Println("对不起，不支持改方法！")
}

func (self *TextFile)getChild(i int) AbstractFile {
	fmt.Println("对不起，不支持该方法！")
	return nil
}

func (self *TextFile) killVirus(){
	fmt.Println("对文本文件",self.name,"进行杀毒")
}

//______________________________________________________//

type Floder struct {
	name string
	fileList []AbstractFile
}

func (self *Floder) Floder(name string) {
	self.name = name
}

func (self *Floder) add(file AbstractFile) {
	self.fileList = append(self.fileList,file)
}

func (self *Floder) remove(file AbstractFile) {
	fileListLen := len(self.fileList)
	for k, v:= range  self.fileList {
		if v == file && k != fileListLen -1  {
			self.fileList=append(self.fileList[:k-1],self.fileList[k+1:]...)
			break
		}
		if  v == file && k == fileListLen -1 {
			self.fileList =self.fileList[:(fileListLen-2)]
		}
	}
}

func (self *Floder) getChild(i int) AbstractFile {
	return self.fileList[i]
}

func (self *Floder) killVirus() {
	fmt.Println("对文件夹",self.name,"进行杀毒")

	for _,v := range self.fileList { //递归,调用成员killVirus方法
		v.killVirus()
	}
}







