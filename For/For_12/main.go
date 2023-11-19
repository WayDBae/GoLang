package main

import "fmt"

func main() {
	var n int
	var a, result float64
	a = 1.1
	result = 1.0
	fmt.Print("N =")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		result *= a
		a += 0.1
	}
	fmt.Println(result)
}
