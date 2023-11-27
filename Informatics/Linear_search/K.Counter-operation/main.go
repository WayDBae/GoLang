package main

import (
	"fmt"
)

func main() {
	var N, max, min int

	fmt.Println("Введите количество элементов массива: ")
	fmt.Scan(&N)
	array := make([]int, N)
	fmt.Println("Введите оценки: ")
	for i := 0; i < N; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println("Журнал создан!")
	fmt.Println("Журнал без замены: ", array)
	max, min = array[0], array[0]
	for i := 0; i < N; i++ {
		if array[i] > max {
			max = array[i]
		} else if array[i] < min {
			min = array[i]
		}
	}
	for i := 0; i < N; i++ {
		if array[i] == max {
			array[i] = min
		}
	}
	fmt.Print("Журнал с заменой: ", array)
}
