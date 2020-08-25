package main

import (
	"fmt"
	"time"
)

// 6) Заданы три числа - a, b, c  - длины сторон треугольника.
// Нужно проверить, является ли треугольник прямоугольным.
// Если является, вывести "Да". Иначе вывести "Нет"
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
