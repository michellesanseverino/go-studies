package main

import "fmt"

func main() {
	fmt.Println("O valor da soma é " + fmt.Sprint(soma(64, 36)))
	fmt.Println("O valor da multiplicação é " + fmt.Sprint(multiplicacao(64, 36)))
	fmt.Prin
}

func soma (x int, y int) int {
	return x + y
}

func multiplicacao (x int, y int ) int {
	return x * y
}