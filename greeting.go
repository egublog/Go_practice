package main
import (
	"fmt"
	"os"
)

type User struct {
	gender string
	age int
	address string
	weight int
}

func main() {
	u := User{"male", 22, "東京都八王子市", 65}

	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(os.Args[0:])

	fmt.Println(u)
	fmt.Println(u.gender)

	fmt.Println(hello())
	fmt.Println(double(19))
}

func hello() (hello string) {
	hello = "hello"
	return
}

func double(number int)(doubleNumber int) {
	doubleNumber = number * 2
	return 
}
