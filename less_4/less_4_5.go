package main

import (
	"fmt"
	"time"
)

// 5) Дано трехзначное число. Переверните его, а затем выведите.
// 1 бал
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
