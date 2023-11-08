package main

import "fmt"

func main() {
	//Даны длины ребер a, b, c прямоугольного параллелепипеда. Найти его объем V = a·b·c и площадь поверхности S = 2·(a·b + b·c + a·c).
	var a, b, c, V, S float64
	fmt.Println("Введите длину ребер a: ")
	fmt.Scan(&a)
	fmt.Println("Введите длину ребер b: ")
	fmt.Scan(&b)
	fmt.Println("Введите длину ребер c: ")
	fmt.Scan(&c)
	V = a * b * c
	S = 2 * (a*b + b*c + a*c)
	fmt.Printf("Объём прямоугольного параллелепипеда = %v\nПлощадь поверхности = %v", V, S)
}
