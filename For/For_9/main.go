package main

import (
	"fmt"
	"math"
)
func main() {
	var a, b, sum int
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println("Ошибка: A должно быть меньше B")
		return
	}
	for i := a; i <= b; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	fmt.Println("Сумма квадратов всех чисел от", a, "до", b, "включительно:", sum)
}
