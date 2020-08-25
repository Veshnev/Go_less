package main

import (
	"fmt"
	"time"
)

// 14) Дан массив из N целых чисел. Выведите +, если массив является красивой горой, иначе  -.
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
