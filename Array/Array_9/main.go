package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	array := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}
	summary := 0
	for i := n-1; i >= 0; i-- {
		if array[i]%2 == 0 {
			fmt.Print(array[i], " ")
			summary++
		}
	}
	fmt.Println()
	fmt.Println(summary)
}
