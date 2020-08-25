package main

import (
	"fmt"
	"time"
)

// 13) На вход подается 2 числа a, b. Из числа a удалить заданную цифру b.
// 2 бала
func main() {
	var n float64

	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	start := time.Now()

	fmt.Println(n)

	fmt.Println("--------------------------")
	start = time.Now()
	fmt.Println("Время выполнения:", time.Since(start))
}
