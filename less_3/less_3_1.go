package main

import (
	"fmt"
	"os"
)

// 1  Подается число N;  Нужно вывести сумму арифметической прогресии числа N; Скорость O(1)
func main() {
	var n uint64

	fmt.Print("Введите число для расчёта арифметической прогрессии: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Zero и в африке — Zero...")
		os.Exit(0)
	}

	fmt.Println("Результат —", (1+n)/2*n+(1+n)%2*n/2)
}
