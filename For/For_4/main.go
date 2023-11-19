package main

import "fmt"

func main() {
	var price, cost float64
	fmt.Println("Введите цену за 1 кг конфет: ")
	fmt.Scan(&price)
	for i := 1; i <= 10; i++ {
		cost = price * float64(i)
		fmt.Printf("%d кг конфет стоит %.2f сомони\n", i, cost)
	}
}
