package main

import "fmt"

func main() {
	number := 532
	ones := number % 10
	tens := (number / 10) % 10
	fmt.Println("Последняя цифра (единицы):", ones)
	fmt.Println("Средняя цифра (десятки):", tens)
}



