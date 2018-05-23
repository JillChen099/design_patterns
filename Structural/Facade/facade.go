/*
Created on 2018/5/21 16:27

author: ChenJinLong

Content: 
*/
package facade

import "fmt"

type FileReader struct {

}

func (f *FileReader) Read(fileName string) string {
	var reStr string
	switch fileName {
	case "A":
		reStr = "第一条明文"
	case "B":
		reStr =  "第二条明文"
	default:
		reStr =  "第三条明文"
	}
	fmt.Println("读取文件，获取明文:",reStr)
	return reStr
}


type CipherMachine struct {

}

func (c *CipherMachine) Encrpty(plainText string) string {

	plainTextByte := []rune(plainText)
	for k,v := range plainTextByte {
		vt := v-2
		plainTextByte[k] = vt
	}
	fmt.Println("数据加密,将明文转换成密文:",string(plainTextByte))
	return string(plainTextByte)
}

type FileWriter struct {

}

func (f *FileWriter) Write(encryStr string,fileNameDes string) {
	fmt.Println("保存密文，写入文件。")
}



//外观模式
type EncryptFacade struct {
	reader *FileReader
	cipher *CipherMachine
	writer *FileWriter
}

func (e *EncryptFacade) FileEncrypt(fileNameSrc,fileNameDes string) {
	plainStr := e.reader.Read(fileNameSrc)
	encryptStr := e.cipher.Encrpty(plainStr)
	e.writer.Write(encryptStr,fileNameDes)
}



