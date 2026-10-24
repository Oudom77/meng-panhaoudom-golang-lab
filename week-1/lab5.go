package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func runLab5(word string) {

	data := []byte(word)

	fmt.Printf("\n\n=== Lab 5 ===")


	hexString := hex.EncodeToString(data)
	fmt.Printf("\nHexadecimal String: %s", hexString)

	base64String := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("\nBase64 String: %s\n", base64String)

	for _, value := range data {

		original := string(value)
		fmt.Printf("%v : %08b\n", original, value)

	}
}