package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	if a < b {
		fmt.Println("Наименьшее число под номером: ",1)
	} else if a > b{
		fmt.Println("Наименьшее число под номером: ",2)
	} else {
		fmt.Println("Оба равны")
	}
}