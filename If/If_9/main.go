package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	if a > b {
		a, b = b, a
	}
	fmt.Printf("Новое значение переменной А = %v\nНовое значение переменной B = %v", a, b)
}
