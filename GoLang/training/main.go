package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Здарова ебать")
	fmt.Println("Введи: \n1 - Калькулятор\n2 - Число умножается на 2 и прибавляется на 100\n3 - Сумма квадратов двух чисел\n4 - Квадрат числа\n5 - Вывод десятка числа\n0 - Выход")
	var choice int
	fmt.Scan(&choice)
L:
	for choice != 0 {
		switch choice {
		case 0:
			break L
		case 1:
			calc()
		case 2:
			task1()
		case 3:
			task2()
		case 4:
			task3()
		case 5:
			task4()
		default:
			fmt.Println("Вы ввели неверное значение!")
		}
	}
}
func calc() {
	var num1 float64
	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)

	var num2 float64
	fmt.Println("Введите второе число: ")
	fmt.Scan(&num2)

	var operator string
	fmt.Print("Введите знак операции (+) (-) (/) (*): ")
	fmt.Scan(&operator)

L:
	switch operator {
	case "+":
		fmt.Println("Ваш ответ: ", num1+num2)
	case "-":
		fmt.Println("Ваш ответ: ", num1-num2)
	case "/":
		if num2 == 0.0 {
			fmt.Println("Divide by Zero Situation")
		} else {
			fmt.Println("Ваш ответ: ", num1/num2)
		}
	case "*":
		fmt.Println("Ваш ответ: ", num1*num2)
	case "!":
		break L
	default:
		fmt.Println("Слыш. Ты не туда нажал.")
	}
}
func task1() {
	/*	Напишите программу, которая последовательно делает следующие операции с введённым числом:
																			Число умножается на 2;
																			Затем к числу прибавляется 100.
	После этого должен быть вывод получившегося числа на экран.*/
	var num int
	fmt.Println("Введите число: ")
	fmt.Scan(&num)
	num = num*2 + 100
	fmt.Println("Ваше конечное число: ", num)
}
func task2() {
	/*	 Петя торопился в школу и неправильно написал программу,
		 которая сначала находит квадраты двух чисел, а затем их суммирует.
		 Исправьте его программу.*/
	var a, b, c int
	fmt.Println("Введите первое число: ")
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Println("Введите второе число: ")
	fmt.Scan(&b) // считаем переменную 'b' с консоли

	a = a * a
	b = b * b
	c = a + b
	fmt.Println("Сумма двух квадратов чисел = ", c)
}
func task3() {
	/*	По данному целому числу, найдите его квадрат.	Формат входных данных				Формат выходных данных
		На вход дается одно целое число.	Программа должна вывести квадрат данного числа.*/
	fmt.Println("Введите число для преобразования в квадрат: ")
	var num int
	fmt.Scan(&num)
	num = num * num
	fmt.Println("Ваше число в квадрате = ", num)
}
func task4() {
	/*Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).

	Формат входных данных
	На вход дается натуральное число, не превосходящее 10000.

	Формат выходных данных
	Выведите одно целое число - число десятков.*/
	var num int
	fmt.Println("Введите число не выше 10К:")
	fmt.Scan(&num)
	if num < 100 {
		s := strconv.Itoa(num)
		fmt.Println(s[0:1])
	} else if num < 1000 {
	s := strconv.Itoa(num)
	fmt.Println(s[1:2])
	} else if num < 10000 {
	s := strconv.Itoa(num)
	fmt.Println(s[2:3])
	} else if num < 0 {
		fmt.Println("Число должно быть положительным!")
	} else {
		fmt.Println("Ты попутал!")
	}
}
	