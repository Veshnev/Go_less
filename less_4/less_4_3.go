package main

import (
	"fmt"
	"time"
)

// 3) На вход подается 5 чисел a, b, c, d, e
// x = a умножается на 2 b раз
// y = c умножается на 2 d раз
// сумма результата x и y делется на 2 e раз сделать сопоставимо по скорости
// 3 бала
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
