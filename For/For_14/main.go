package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Введите целое число N (> 0): ")
	fmt.Scan(&N)

	sum := 0
	output := ""

	for i := 1; i <= 2*N-1; i += 2 {
		sum += i
		output += fmt.Sprintf("%d ", sum)
	}

	fmt.Printf("Квадрат числа %d: %d\n", N, sum)
	fmt.Printf("Промежуточные значения: %s\n", output)
}
