package main

import (
	"fmt"
	"time"
)

// 11) Дано число n. Определите, каким по счету числом Фибоначчи оно является, Если n не является числом Фибоначчи, выведите число 0.
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
