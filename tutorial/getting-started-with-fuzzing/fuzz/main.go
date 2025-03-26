package main

import "fmt"

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
 
func Reverse(s string) string {
	r := []rune(s)
	for i := range len(r)/2 {
		j := len(r) - i - 1
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}