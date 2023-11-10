package main

import "fmt"

func main() {
	/*Найти длину окружности L и площадь круга S заданного радиуса R:
	L = 2·π·R, S = π·R^2.
	В качестве значения π использовать 3.14.*/
	var L, S, R, pi float64
	fmt.Println("Введите радиус окружности R: ")
	fmt.Scan(&R)
	pi = 3.14
	L = 2 * pi * R
	S = pi * (R * R)
	fmt.Printf("Длина окружности L = %v\nПлощадь круга S = %v", L, S)
}
