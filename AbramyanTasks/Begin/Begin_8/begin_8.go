package main

import "fmt"

func main() {
	//Даны два числа a и b. Найти их среднее арифметическое: (a + b)/2.
	var a, b, average float64
	fmt.Println("Введите а: ")
	fmt.Scan(&a)
	fmt.Println("Введите b: ")
	fmt.Scan(&b)
	average = (a + b) / 2
	fmt.Printf("Среднеарифметическое число чисел %v и %v равно %v", a, b, average)
}
