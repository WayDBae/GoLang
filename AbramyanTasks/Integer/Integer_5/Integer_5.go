package main

import "fmt"

func main() {
	A := 17
	B := 5
	unused := A % B
	fmt.Println("Длина незанятой части отрезка A:", unused)
}
