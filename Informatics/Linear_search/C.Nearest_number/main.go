package main

import (
	"fmt"
	"math"
)

func main() {
	var N, x int

	fmt.Println("Введите количество элементов массива: ")
	fmt.Scan(&N)
	array := make([]int, N)
	fmt.Println("Введите элементы массива: ")
	for i := 0; i < N; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println("Массив создан")

	closest := array[0]
	fmt.Println("Какое число хотите проверить?")
	fmt.Scan(&x)
	for i := 0; i < N; i++ {
		if math.Abs(float64(array[i]-x)) < math.Abs(float64(closest-x)) {
			closest = array[i]
		}
	}
	fmt.Println(closest)
}
