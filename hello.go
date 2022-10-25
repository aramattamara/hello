package main

import "log"

//	func main() {
//		fmt.Println("Hello, World!")
//	}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main3() {
	x, y := split(17)
	log.Printf("%v, %v", x, y)
}
