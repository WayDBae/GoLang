package main

import "fmt"

func main() {
	var a,b,c,positive,negative int
	fmt.Scan(&a, &b, &c)
	if a > 0 {positive++} else if a < 0 {negative--}
	if b > 0 {positive++} else if b < 0 {negative--}
	if c > 0 {positive++} else if c < 0 {negative--}
	fmt.Printf("Количество положительных чисел в исходном наборе: %v\nКоличество отрицательных чисел в исходном наборе: %v",positive,negative)
}
