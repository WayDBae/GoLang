package main

import "fmt"

func main() {
	//Дан диаметр окружности d. Найти ее длину L = π·d. В качестве значения π использовать 3.14
	var d, L, pi float64
	pi = 3.14
	fmt.Println("Введите диаметр окружности: ")
	fmt.Scan(&d)
	L = pi * d
	fmt.Println("Длина окружности равна: ", L)
}
