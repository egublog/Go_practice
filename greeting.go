package main
import ("fmt")

type User struct {
	gender string
	age int
}

func main() {
	u := User{"male", 22}

	fmt.Println(u)
}
