// https://blog.carlmjohnson.net/post/google-go-the-good-the-bad-and-the-meh/

package main

import "fmt"

func accgen(n int) func(int) int {
	return func(i int) int {
		n += i
		return n
	}
}

func main() {
	f := accgen(0)
	fmt.Println(f(1), f(1), f(1)) //Prints 1, 2, 3
}
