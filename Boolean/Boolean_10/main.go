package main

import "fmt"

func main() {
	var a, b int
	var result bool
	fmt.Scan(&a, &b)
	result = ((a%2 != 0) && (b%2 == 0)) || ((a%2 == 0 && b%2 != 0))
	fmt.Println(result)
}