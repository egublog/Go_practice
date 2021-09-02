package main
import (
	"fmt"
	"os"
)

type User struct {
	gender string
	age int
	address string
}

func main() {
	u := User{"male", 22, "東京都八王子市"}

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Println(u)
	fmt.Println(u.gender)

	fmt.Println(hello())
}

func hello() (hello string) {
	hello = "hello"
	return
}
