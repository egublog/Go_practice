package main

import (
	"fmt"
)

func goroutine(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go goroutine("world")
	normal("hello")
}

func circleArea(pai float64) func(radius float64) float64 {
	return func (radius float64) float64 {
		return pai * radius * radius
	}
}
