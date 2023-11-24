package main

import "fmt"

func main() {
	var N int
	fmt.Print("Введите размер массива: ")
	fmt.Scan(&N)

	array := make([]int, N)

	fmt.Println("Введите элементы массива:")
	for i := 0; i < N; i++ {
		fmt.Scan(&array[i])
	}

	fmt.Println("Массив в обратном порядке:")
	for i := N - 1; i >= 0; i-- {
		fmt.Print(array[i], " ")
	}
}	