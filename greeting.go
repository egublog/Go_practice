package main

import "fmt"


type User struct {
	name string
	age int
}

func (u *User) action() {
	u.name = u.name + "a"
	u.age  = u.age * 3
}

func main() {
	user := User{"江口", 23}
	user.action()
	fmt.Println(user)
}

func circleArea(pai float64) func(radius float64) float64 {
	return func (radius float64) float64 {
		return pai * radius * radius
	}
}
