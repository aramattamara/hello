package main

import "fmt"

//	func main() {
//		fmt.Println("Hello, World!")
//	}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
