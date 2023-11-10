package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Println(a)
	} else {fmt.Println(b)}
}