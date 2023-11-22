package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	quotient := 0
	remainder := N

	for remainder >= K {
		remainder = remainder - K
		quotient++
	}

	fmt.Printf("Частное: %d, Остаток: %d\n", quotient, remainder)
}
