package main

import "fmt"

// На ввод подается число секунд,  нужно вывести сколько это часов, минут и секунд
func main() {
	var a int

	fmt.Scan(&a)
	fmt.Println(a % 100 / 10)
}
