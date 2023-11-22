package main

import (
	"fmt"
)

func main() {
	var A,B int
	fmt.Scan(&A, &B)
	count := 0
	for A >= B {
		A -= B
		count++
	}
	fmt.Println(count) // Вывод: 5
}
