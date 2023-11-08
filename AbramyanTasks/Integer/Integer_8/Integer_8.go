package main

import "fmt"

func main() {
	number := 42
	reversed := (number%10)*10 + (number / 10)
	fmt.Println("Число после перестановки цифр:", reversed)
}
