package main

import "fmt"

// На ввод подается число секунд,  нужно вывести сколько это часов, минут и секунд
func main() {
	var a int

	fmt.Scan(&a)

	hours := a / 60 / 60
	minutes := (a - hours * 3600) / 60
	seconds := a % 60

	fmt.Println("Часов: ", hours)
	fmt.Println("Минут: ", minutes)
	fmt.Println("Секунд: ", seconds)
}
