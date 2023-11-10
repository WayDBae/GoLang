package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a > 2 && b <= 3)
}