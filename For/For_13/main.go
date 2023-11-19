package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Введите целое число N (> 0): ")
	fmt.Scan(&N)

	sum := 0.0
	sign := 1.0

	for i := 1; i <= N; i++ {
		term := float64(i) / 10.0 * sign
		sum += term
		sign *= -1.0
	}

	fmt.Printf("Значение выражения: %.4f\n", sum)
}
