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

	sum := 0
	for i := n; i <= 2*n; i++ {
		sum += i * i
	}

	fmt.Println("Сумма ряда N^2 + (N+1)^2 + (N+2)^2 + ... + (2N)^2:", sum)
}
