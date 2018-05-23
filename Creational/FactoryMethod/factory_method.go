/*
Created on 2018/5/17 16:29

author: ChenJinLong

Content:
工厂方法模式(Factory Method Pattern)：定义一个用于创建对象的接口，让子类决定将哪一个类实例化。
在工厂方法模式中，我们不再提供一个统一的工厂类来创建所有的产品对象，
而是针对不同的产品提供不同的工厂，系统提供一个与产品等级结构对应的工厂等级结构。
*/
package FactoryMethod

import "fmt"

//产品接口
type Logger interface {
	writeLog()

}

//具体产品
type DatabaseLogger struct {

}
func (d *DatabaseLogger) writeLog() {
	fmt.Println("数据库日志记录")
}

type FileLogger struct {

}

func (f *FileLogger) writeLog() {
	fmt.Println("文件日志记录")
}



//工厂接口
type LoggerFactory interface {
	createLogger() Logger

}


//具体产品工厂
type DatabaseLoggerFactory struct {

}

func (d *DatabaseLoggerFactory) createLogger() Logger {
	logger := new(DatabaseLogger)
	return logger
}


type FileLoggerFactory struct {

}

func (f *FileLoggerFactory) createLogger() Logger {
	logger := new(FileLogger)
	return logger
}