package main

import "fmt"

func main() {
	var a, b int
	n := 0
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println("Ошибка: A должно быть меньше B")
		return
	}
	for i := a; i <= b; i++ {
		fmt.Println(i)
		n+=1
	}
	fmt.Println(n)
}
