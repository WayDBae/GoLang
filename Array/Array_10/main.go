package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	array := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	for i := 0; i < n; i++ {
		if array[i]%2 == 0 {
			fmt.Print(array[i], " ")
		}
	}

	for i := n - 1; i >= 0; i-- {
		if array[i]%2 != 0 {
			fmt.Print(array[i], " ")
		}
	}
}
