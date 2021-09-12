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

	a := circleArea(3.14)
	fmt.Println(a(3))
}

func circleArea(pai float64) func(radius float64) float64 {
	return func (radius float64) float64 {
		return pai * radius * radius
	}
}
