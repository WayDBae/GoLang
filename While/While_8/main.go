package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	K := 1 // Начальное значение K

	for K*K <= N {
		K++
	}

	K--
	
	fmt.Println(K)
}
