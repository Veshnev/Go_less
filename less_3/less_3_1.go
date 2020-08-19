package main

import (
	"fmt"
	"os"
)

// 1  Подается число N;  Нужно вывести сумму арифметической прогресии числа N; Скорость O(1)
func main() {
	var N uint64

	fmt.Print("Введите число для расчёта арифметической прогрессии: ")
	fmt.Scan(&N)

	if N == 0 {
		fmt.Println("Zero и в африке — Zero...")
		os.Exit(0)
	}

	fmt.Println("Результат: ", (1+N)/2*N+(1+N)%2*N/2)
}
