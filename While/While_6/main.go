package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	result := 1.0

	for N >= 2 {
		result *= float64(N)
		N -= 2
	}

	fmt.Println(result)
}