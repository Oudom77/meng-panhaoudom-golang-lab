package main

import "fmt"

func runLab3(n1 int, n2 int) {

	fmt.Printf("\n\n=== Lab 3 ===\n")

	myAnd(n1, n2)
	myOr(n1, n2)
	myXor(n1, n2)
	myNot(n1, n2)
	myLeftShift(n1, n2)
	myRightShift(n1, n2)

}

func myXor(n1 int, n2 int) {

	fmt.Printf("\n%d XOR %d is %b\n",n1, n2, n1 ^ n2)

}

func myNot(n1 int, n2 int) {

	fmt.Printf("%d NOT is %b, %d NOT is %b\n",n1, ^n1, n2, ^n2)

}

func myLeftShift(n1 int, n2 int) {

	fmt.Printf("%d LeftShift %d is %b\n",n1, n2, n1 << n2)

}

func myRightShift(n1 int, n2 int) {

	fmt.Printf("%d RightShift %d is %b\n",n1, n2, n1 >> n2)

}

func myAnd(n1 int, n2 int) {

	fmt.Printf("%d AND %d is %b\n",n1, n2, n1 & n2)

}

func myOr(n1 int, n2 int) {

	fmt.Printf("%d OR %d is %b\n",n1, n2, n1 | n2)

}