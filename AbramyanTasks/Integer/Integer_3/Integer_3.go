package main

import "fmt"

func main() {
	fileSize := 15000
	kilobytes := fileSize / 1024
	fmt.Println("Количество полных килобайтов:", kilobytes)
}
