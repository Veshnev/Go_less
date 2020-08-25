package main

import (
	"fmt"
	"time"
)

// 7) Даны три числа a, b, c. Определите, существует ли треугольник с такими сторонами.
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
