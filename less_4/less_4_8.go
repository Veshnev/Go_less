package main

import (
	"fmt"
	"time"
)

// 8) Вводится одно число n, Вывести цифровой корень числа n.
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
