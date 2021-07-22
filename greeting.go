package main
import ("fmt")

func main() {
	num := 3
	fmt.Println(num)
	fmt.Println("Good morning")
	fmt.Println("Good afternoon")
	fmt.Println("Good evening")
	array()
}

func array() {
	a := [...]string{"a", "b", "c"}
	fmt.Println(a)
}
