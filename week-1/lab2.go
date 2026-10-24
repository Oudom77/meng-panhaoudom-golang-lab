package main

import "fmt"

func runlab2(n1 int, n2 int) {

	fmt.Printf("\n\n=== Lab 2 ===\n")

	if (n1 > 0 && n2 > 0){

		fmt.Printf("Positive: True\n")

	} else {

		fmt.Printf("Positive: False\n")

	}

	if (n1 > n2 || n2 > n1){

		fmt.Printf("One greater than the other: True\n")

	} else {

		fmt.Printf("One greater than the other: False\n")

	}

	if (n1 != n2){

		fmt.Printf("Not equal: True\n")

	} else {

		fmt.Printf("Not equal: False\n")

	}

}