package main

import (
	"fmt"
	"os"
)

// 2) Подается число N;  и подется ...X числе N
// 		3
// 		52 64 108
// 	  сумму двузначных чисел X , кратных 8
func main() {
	var N, X, result int

	fmt.Print("Введите целевое количество чисел: ")
	fmt.Scan(&N)

	if N == 0 {
		fmt.Println("Филосовсий подход... Спасибо за внимание, пока! ))")
		os.Exit(0)
	}

	fmt.Println("--------------------")

	for i := 1; i <= N; i++ {
		if i == N {
			fmt.Print("Введите последнее число: ")
		} else {
			fmt.Print("Введите число ", i, ": ")
		}

		fmt.Scan(&X)

		if X%8 == 0 {
			result += X
		}
	}

	fmt.Println("--------------------")
	fmt.Println("Сумма чисел, делящихся на 8: ", result)
}
