package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N)
	K = 0
	for N%2 == 0 {
		N /= 2
		K++
	}

	fmt.Println(K)
}