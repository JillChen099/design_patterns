/*
Created on 2018/5/28 11:43

author: ChenJinLong

Content:
客户端不想或不能直接访问一个对象，
此时可以通过一个称之为“代理”的第三者来实现间接访问，
该方案对应的设计模式被称为代理模式。
*/
package proxy

import "fmt"

type AccessValidator struct {

}

func (self *AccessValidator) Validate(userId string) bool {
	fmt.Printf("在数据库中验证用户%s是否是合法用户？\n",userId)
	if userId == "杨过" {
		fmt.Printf("%s登录成功！\n",userId)
		return true
	}else {
		fmt.Printf("%s登录失败！\n",userId)
		return false
	}
}

type Logger struct {

}

func (self *Logger)Log(userId string) {

	fmt.Printf("更新数据库，用户%s查询次数+1！\n",userId)
}

//抽象主题角色
type Searcher interface {
	DoSearch(userId,keyword string) string

}

//真实主题角色
type RealSearch struct {

}

func (self *RealSearch) DoSearch(userId ,keyword string) string {
	fmt.Printf("用户%s使用关键词%s查询商务信息\n",userId,keyword)
	return "返回具体内容"
}

//代理主题角色
type ProxySearcher struct {
	searcher  *RealSearch
	validator  *AccessValidator
	logger    *Logger
}

func (self *ProxySearcher) DoSearch(userId,keyword string) string {
	if self.validator.Validate(userId) {
		result := self.searcher.DoSearch(userId,keyword)
		self.logger.Log(userId)
		return result
	}else {
		return ""
	}

}




