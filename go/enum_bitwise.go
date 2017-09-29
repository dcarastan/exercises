// https://blog.carlmjohnson.net/post/google-go-the-good-the-bad-and-the-meh/

package main

import "fmt"

type BFlag uint8

const (
	Grapes BFlag = 1 << iota
	Apples
	Bananas
	Mangos
)

func OrderBreakfast(flag BFlag) {
	if flag&Apples != 0 {
		fmt.Println("Apples it is!")
	} else {
		fmt.Println("All we had was apples.")
	}
}

func main() {
	f := Grapes | Apples | Bananas
	OrderBreakfast(f)
}
