package main

import "fmt"

func main() {
	//Задание 1
	fmt.Printf("Привет Хумо!\n\n")

	//Задание 2
	var num1, num2 int
	num1 = 10
	num2 = 3
	fmt.Println("Сложение num1 + num2: ", num1+num2)
	fmt.Println("Вычитание num1 - num2: ", num1-num2)
	fmt.Println("Деление num1 / num2: ", num1/num2)
	fmt.Println("Умножение num1 * num2: ", num1*num2)
	fmt.Println("Деление с остатком num1 % num2: ", num1%num2)

	//Задание 3
	var a int
	var b float32
	var c bool
	var d complex64
	var name string
	fmt.Println(a, b, c, d, name)
}
