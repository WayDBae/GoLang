package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a != b {
		a += b
	} else {
		a, b = 0, 0
	}
	fmt.Println(a,b)
}
