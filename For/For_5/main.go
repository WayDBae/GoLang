package main

import "fmt"

func main() {
	var price float64
	fmt.Println("Введите цену за 1 кг конфет: ")
	fmt.Scan(&price)
	for i := 1; i <= 10; i++ {
        weight := float64(i) / 10.0
    	cost := price * weight
        fmt.Printf("%.1f кг конфет стоит %.2f\n", weight, cost)
    }
}
