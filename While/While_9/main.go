package main

import "fmt"

func main() {
	var N, K int

	fmt.Print("Введите целое число N: ")
	fmt.Scanln(&N)

	K = 1
	for 3*K <= N {
		K++
	}

	fmt.Println("Наименьшее целое число K:", K)
}