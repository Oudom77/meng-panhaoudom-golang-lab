package main

import "fmt"

func runlab1(n1 int, n2 int) int {

	var result int = n1
	var choice int

	fmt.Println("=== Lab 1 ===")


	fmt.Println("Choose an operator: ")
	fmt.Println("1. +")
	fmt.Println("2. -")
	fmt.Println("3. *")
	fmt.Println("4. /")
	fmt.Println("5. %")

	fmt.Printf("Choice: ")
	fmt.Scan(&choice)

	switch choice {
		case 1:
			result += n2
		case 2:
			result -= n2
		case 3:
			result *= n2
		case 4:
			result /= n2
		case 5:
			result %= n2
	}

	return result
}