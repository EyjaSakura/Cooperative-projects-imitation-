package main

import "fmt"

var name, sex string
var age int

func getInfo() {
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Print("请输入性别：")
	fmt.Scanln(&sex)
}

func reInfo() (string, int, string) {
	return name, age, sex
}

func main() {
	getInfo()
	username, _, _ := reInfo()
	fmt.Println("你输入的姓名是：" + username)
}
