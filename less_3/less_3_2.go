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
	var n, d, x, result int

	fmt.Print("Введите целевое количество чисел: ")
	fmt.Scan(&n)

	fmt.Print("Введите делитель: ")
	fmt.Scan(&d)

	if n == 0 || d == 0 {
		fmt.Println("Филосовсий подход... Спасибо за внимание, пока! ))")
		os.Exit(0)
	}

	fmt.Println("--------------------")

	for i := 1; i <= n; i++ {
		if i == n {
			fmt.Print("Введите последнее число: ")
		} else {
			fmt.Print("Введите число ", i, ": ")
		}

		fmt.Scan(&x)

		if x%d == 0 {
			result += x
		}
	}

	fmt.Println("--------------------")
	fmt.Println("Сумма чисел, делящихся на", d, "—", result)
}
