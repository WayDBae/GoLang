package main

import (
	"fmt"
)

//Дана сторона квадрата a. Найти его площадь S = a^2.
func main() {
	var a, square float64
	fmt.Println("Введите сторону квадрата а: ")
	fmt.Scan(&a)
	square = a * a
	fmt.Println("Площадь квадрата равна: ", square)
}
