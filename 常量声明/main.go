package main

import "fmt"

func main() {
	/*
			1、常量的定义方法：const identifier [type] = value
			2、枚举类型
			定义性别枚举
			const (
		    Unknown = 0
		    Female  = 1
		    Male    = 2
			)
			3、常量表达式中可以使用 len()、cap()、unsafe.Sizeof() 等内置函数计算表达式的值。
			4、iota 是 Go 语言的特殊常量，也称为常量计数器。它可以在 const 声明块中自动生成递增的序列值。
	*/
	// iota 从 0 开始，每行加 1
	const (
		a = iota // 0
		b        // 1（省略值，默认使用上一行的 iota）
		c        // 2
	)
	fmt.Println(a, b, c) //0 1 2

}
