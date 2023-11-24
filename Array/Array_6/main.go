package main

import "fmt"

func main() {
	var N, A, B int
	fmt.Print("Введите значение N: ")
	fmt.Scan(&N)
	fmt.Print("Введите значение A: ")
	fmt.Scan(&A)
	fmt.Print("Введите значение B: ")
	fmt.Scan(&B)

	array := make([]int, N)
	array[0] = A
	array[1] = B

	for i := 2; i < N; i++ {
		array[i] = array[i-1] + array[i-2]
	}

	fmt.Println(array)
}