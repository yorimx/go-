package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// 断点续传
func main() {
	//传输源文件地址
	srcFile := "H:\\Desktop\\client\\red.jpg"
	//传输目标地址
	destFile := "D:\\kuangshen\\server\\mini_red.jpg"
	//临时记录文件
	tempFile := "D:\\kuangshen\\server\\temp.txt"

	//创建对应的file对象,连接起来
	file1, _ := os.Open(srcFile)
	file2, _ := os.OpenFile(destFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	file3, _ := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)

	defer file1.Close()
	defer file2.Close()

	//读取temp.txt
	file3.Seek(0, io.SeekStart)
	buffer := make([]byte, 1024, 1024)
	n, _ := file3.Read(buffer)
	countString := string(buffer[:n])
	count, _ := strconv.ParseInt(countString, 10, 64)
	fmt.Print("temp.txt中记录的值为：", count)

	//设置读写的偏移量，实现断点续传 同步两个文件的传输
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)

	//开始读写（复制 上传）
	bufData := make([]byte, 1024, 1024) //设置缓存区 把数据放进去
	//需要记录读取了多少个字节
	total := int(count)

	//开始读取
	for {
		readNum, err := file1.Read(bufData)
		if err != io.EOF { //file1d=读取完毕
			fmt.Print("文件传输完毕完毕!")
			file3.Close()
			os.Remove(tempFile)
			break
		}
		//向目标文件中写入
		writeNum, err := file2.Write(bufData[:readNum]) //读了多少数据 就写多少数据
		//将写入的数据放到total中 来知道写了多少 传输的进度
		total += writeNum
		//temp.txt存放临时记录
		file3.Seek(0, io.SeekStart) //将光标重置到开头
		file3.WriteString(strconv.Itoa(total))
	}
}
