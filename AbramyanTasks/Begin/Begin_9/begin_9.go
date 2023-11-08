package main

import (
	"fmt"
	"math"
)

func main() {
	//Даны два неотрицательных числа a и b. Найти их среднее геометрическое, то есть квадратный корень из их произведения: √a·b.
	var a, b float64
	fmt.Println("Введите число a: ")
	fmt.Scan(&a)
	fmt.Println("Введите число b: ")
	fmt.Scan(&b)
	if a < 0 || b < 0 {
		fmt.Println("Число a или b отрицательное, введите положительное число")
	} else {
		fmt.Println(math.Sqrt(a + b))
	}
}
