package main

import "fmt"

func main() {
	var N, K int

	fmt.Print("Введите целое число N: ")
	fmt.Scanln(&N)

	K = 1
	for 3*K < N {
		K++
	}

	// Если условие неравенства никогда не выполнялось, уменьшаем значение K на 1
	if 3*K >= N {
		K--
	}

	fmt.Println("Наибольшее целое число K:", K)
}