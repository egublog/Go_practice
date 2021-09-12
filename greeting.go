package main

import "fmt"


func main() {
	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int
	p1 = &i
	p2 = &j

	i = *p1 + *p2
	fmt.Println(*p1)
	i = *p1 + *p2
	fmt.Println(*p1)
	i = *p1 + *p2
	fmt.Println(*p1)
	i = *p1 + *p2
	fmt.Println(*p1)
	i = *p1 + *p2
	fmt.Println(*p1)
}

func circleArea(pai float64) func(radius float64) float64 {
	return func (radius float64) float64 {
		return pai * radius * radius
	}
}
