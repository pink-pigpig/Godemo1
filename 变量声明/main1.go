package main

import "fmt"

func main() {
	// var name type最简单的使用
	var name string
	fmt.Printf("name: %v\n", name) //如果指定类型后没有赋值，那么值就是默认值
	name = "zhangsan"
	fmt.Printf("name: %v\n", name)
	// %v是指以值的默认格式化输出

	var name1 string = "lisi" //这里其实可以省略掉string,go会自动推导类型
	fmt.Printf("name1: %v\n", name1)
	//var name1 string = "wanwu"错误的。声明不要重复声明,go不支持同一作用域重复声明

	var name2 = "lisi"
	fmt.Printf("name2: %v\n", name2)

	var name3, name4 string //支持多个声明变量
	name3 = "lisi"
	name4 = "lisi"
	fmt.Printf("name3: %v, name4: %v\n", name3, name4)

	//支持多个声明变量,不同类型。
	var (
		name5  string
		age    int
		height float64
		isMale bool
	)
	name5 = "lisi"
	age = 18
	height = 1.78
	isMale = true
	fmt.Printf("name5: %v, age: %v, height: %v, isMale: %v\n", name5, age, height, isMale)

	//支持多个声明变量,不同类型。支持类型推导。
	var (
		name6   = "lisi"
		age1    = 18
		height1 = 1.78
		isMale1 = true
	)

	fmt.Printf("name6: %v, age1: %v, height1: %v, isMale1: %v\n", name6, age1, height1, isMale1)

	// 短变量声明法！直接变量命直接赋值,同样支持多个变量
	username := "lisi"
	a, b, c := 1, 2, 3
	fmt.Printf("a: %v, b: %v, c: %v\n", a, b, c)
	fmt.Printf("username: %v\n", username)

	//可以通过%T来获得输出的值是什么类型。
	fmt.Printf("username: %v, username type: %T\n", username, username)

}
