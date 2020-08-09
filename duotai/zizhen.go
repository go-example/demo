package main

import "fmt"

func main() {
	//指针的指针
	pointer()
	//交换变量
	change()
}

func pointer() {
	var a int
	var ptr *int
	var pptr **int
	var ppptr ***int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr
	ppptr = &pptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
	fmt.Printf("指向指针的指针变量 ***ppptr = %d\n", ***ppptr)
}

func change() {
	/* 定义局部变量 */
	var a = 100
	var b = 200

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	/* 调用函数用于交换值
	* &a 指向 a 变量的地址
	* &b 指向 b 变量的地址
	 */
	swap(&a, &b)
	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
