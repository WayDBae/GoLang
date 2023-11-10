package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a > 0 {
		a++
	} else if a < 0 {
		a -= 2
	} else {
		a = 10
	}
	fmt.Println(a)
}
