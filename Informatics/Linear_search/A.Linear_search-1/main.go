package main

import "fmt"

func main() {
	var N, x, result int
	fmt.Println("Введите количество элементов массива: ")
	fmt.Scan(&N)
	array := make([]int, N)
	fmt.Println("Введите элементы массива: ")
	for i := 0; i < N; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println("Массив создан")
	fmt.Println("Какое число хотите проверить?")
	fmt.Scan(&x)
	for i := 0; i < N; i++ {
		if array[i] == x {
			result++
		}
	}
	fmt.Println(result)
}
