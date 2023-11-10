package main

import "fmt"

func main() {
	//Даны два ненулевых числа. Найти сумму, разность, произведение и частное их квадратов
	var a, b float64
	// Ввод чисел
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	// Вычисление
	sum := a*a + b*b
	diff := a*a - b*b
	prod := (a * a) * (b * b)
	quot := (a * a) / (b * b)

	// Вывод результатов
	fmt.Println("Сумма их квадратов:", sum)
	fmt.Println("Разность их квадратов:", diff)
	fmt.Println("Произведение их квадратов:", prod)
	fmt.Println("Частное их квадратов:", quot)
}
