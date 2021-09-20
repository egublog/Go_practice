package main

import (
	"fmt"
	"../Go_practice/mylib"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine1(s, c)
	x := <-c
	fmt.Println(x)

	h := mylib.Average(s)
	
	fmt.Println(h)
}

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
