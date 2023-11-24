package main

import "fmt"

func main() {
	var N int
    fmt.Print("Введите значение N: ")
    fmt.Scan(&N)

    array := make([]int, N)

    for i := 0; i < N; i++ {
        array[i] = 2*i + 1
    }

    fmt.Println(array)
}
