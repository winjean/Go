//demo_24.go
package main

import "fmt"

func main() {
	i := 1
	fmt.Printf("当 i = %d 时：\n", i)

	switch i {
	case 1:
		fmt.Println("输出 i =", 1)
	case 2:
		fmt.Println("输出 i =", 2)
	case 3:
		fmt.Println("输出 i =", 3)
		// 会直接运行紧跟的后一个case或default语句，不论条件是否满足都会执行
		fallthrough
	case 4, 5, 6:
		fmt.Println("输出 i =", "4 or 5 or 6")
	case 7:
		fmt.Println("输出 i =", "7")
	default:
		fmt.Println("输出 i =", "xxx")
	}
}
