package main
import ("fmt")

type User struct {
	gender string
	age int
	address string
}

func main() {
	u := User{"male", 22, "東京都八王子市"}

	a := 439
	b := &a

	c := [...]int {1, 2, 3, 4, 5, 56}
	d := c[1:3]

	fmt.Println(u)
	fmt.Println(u.gender)
	fmt.Println(*b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(hello())
}

func hello() (hello string) {
	hello = "hello"
	return
}
