package main

import "fmt"

func main() {
	//Даны стороны прямоугольника a и b. Найти его площадь S = a·b и периметр P = 2·(a + b).
	var a, b, S, P float64
	fmt.Println("Введите сторону прямоугольника a: ")
	fmt.Scan(&a)
	fmt.Println("Введите сторону прямоугольника b: ")
	fmt.Scan(&b)
	S = a * b
	P = 2 * (a + b)
	fmt.Printf("Плошадь прямоугольника = %v\nПериметр прямоугольника равен = %v", S, P)
}
