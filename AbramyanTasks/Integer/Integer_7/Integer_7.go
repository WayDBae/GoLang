package main

import "fmt"

func main() {
	number := 27
	digit1 := number / 10
	digit2 := number % 10
	sum := digit1 + digit2
	product := digit1 * digit2
	fmt.Println("Сумма цифр:", sum)
	fmt.Println("Произведение цифр:", product)
}
