package main

import "fmt"

func main() {
	number := 56
	tens := number / 10
	ones := number % 10
	fmt.Println("Левая цифра (десятки):", tens)
	fmt.Println("Правая цифра (единицы):", ones)
}
