package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Введите значение N: ")
	fmt.Scan(&N)

	array := make([]int, N)

	for i := 0; i < N; i++ {
		array[i] = int(math.Pow(2, float64(i)))
	}

	fmt.Println(array)
}