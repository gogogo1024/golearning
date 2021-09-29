//go 变量，指针，地址
package main

import (
	"fmt"
)
func main() {
	var house = "yichun"

	// 对字符串取地址
	ptr :=&house

	// 输出ptr的类型
	fmt.Printf("ptr type: %T\n",ptr)

	// 输出ptr指针地址
	fmt.Printf("ptr address: %p\n",ptr)

	// 对指针进行取值操作
	value :=*ptr

	// 输出取值后value的类型
	fmt.Printf("value type: %T\n",value)


	// 输出value的值也就是对指针取值
	fmt.Printf("value: %s\n",value)








}