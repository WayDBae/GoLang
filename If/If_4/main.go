package main

import "fmt"

func main() {
	var a,b,c,sum int
	fmt.Scan(&a, &b, &c)
	if a > 0 {sum++}
	if b > 0 {sum++}
	if c > 0 {sum++}
	fmt.Println("Количество чисел в исходном наборе: ",sum)
}
