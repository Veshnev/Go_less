package main

import "fmt"

// Подается число N; N > 999 && N < 10000; Нужно выыести число десятков
func main() {
	var a int

	fmt.Scan(&a)
	fmt.Println(a % 100 / 10)
}
