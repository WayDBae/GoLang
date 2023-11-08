package main

import "fmt"

func main() {
	//Дана длина ребра куба a. Найти объем куба V = a 3 и площадь его поверхности S = 6·a 2 .
	var a, V, S float64
	fmt.Println("Введите длину ребра куба а: ")
	fmt.Scan(&a)
	V = a * a * a
	S = 6 * (a * a)
	fmt.Printf("Объем куба равен = %v\nПлощадь его поверхности = %v", V, S)
}
