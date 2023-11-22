package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Print("Введите длину отрезка A: ")
	fmt.Scanln(&A)
	fmt.Print("Введите длину отрезка B: ")
	fmt.Scanln(&B)

	remainder := A
	for remainder >= B {
		remainder -= B
	}

	fmt.Println("Длина незанятой части отрезка A:", remainder)
}