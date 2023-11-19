package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Введите значение N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Ошибка: N должно быть больше 0")
		return
	}

	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i)
	}

	fmt.Println("Сумма ряда 1 + 1/2 + 1/3 + ... + 1/N:", sum)
}