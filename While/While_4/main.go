package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	isPowerOfThree := false

	for N > 1 {
		if N%3 != 0 {
			break
		}
		N /= 3
	}

	if N == 1 {
		isPowerOfThree = true
	}

	fmt.Println(isPowerOfThree)
}