package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	hundreds := number / 100
	fmt.Println("Первая цифра (сотни):", hundreds)
}
