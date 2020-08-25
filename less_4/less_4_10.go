package main

import (
	"fmt"
	"time"
)

// 10) Вводится одно число n. n < 10000 Вывести его буквами на русском языке
// 4 бала
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
