package main

import "fmt"

func main() {

    lab1 := runlab1(1, 1)

    fmt.Println("Hello!, Golang")

    fmt.Printf("Lab 1 result: %d", lab1)

    runlab2(1, 1)

    runLab3(2, 2)

    lab4, isError := runlab4(2, 0)

    if (!isError){
        fmt.Printf("\nLab 4 result: %d", lab4)
    }

    runLab5("hi")

    fmt.Printf("\n=== Lab 6 ===\n")


    message := "Secret cookie!"
    lab6Encrypted := xorEncrypt(message, 101)

    fmt.Printf("\nOriginal: %s", message)
    fmt.Printf("\nXOR encryption: %X", lab6Encrypted)

    lab6Decrypted := xorEncrypt(lab6Encrypted, 101)

    fmt.Printf("\nXOR decrypted: %s", lab6Decrypted)
}

