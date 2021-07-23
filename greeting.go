package main
import ("fmt")

type User struct {
	gender string
	age int
	address string
}

func main() {
	u := User{"male", 22, "東京都八王子市"}

	fmt.Println(u)
}
