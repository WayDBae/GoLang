package main

import "fmt"

func main() {
//Дана сторона квадрата a. Найти его периметр P = 4·a.
	var a, perimeter float64
	fmt.Println("Введите сторону квадрата а: ")
	fmt.Scan(&a)
	perimeter = 4 * a
	fmt.Print("Периметр квадрата равен: ", perimeter)
}

