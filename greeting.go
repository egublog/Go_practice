package main

import "fmt"


func main() {
	fmt.Println("aa")
}

func circleArea(pai float64) func(radius float64) float64 {
	return func (radius float64) float64 {
		return pai * radius * radius
	}
}
