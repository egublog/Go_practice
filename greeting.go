package main

import "fmt"


func main() {
	var	c []int

	c = make([]int, 0, 5)
	for i := 1; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}

	func (x int) {
		r := x * 2
		fmt.Println(r)
	}(98)
}
