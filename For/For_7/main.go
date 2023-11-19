package main

import "fmt"

func main() {
	var a, b, count int
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println("Ошибка: A должно быть меньше B")
		return
	}
	count = 0
	for i := a; i <= b; i++ {
		count += i
	}
	fmt.Println(count)
}
