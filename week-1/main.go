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
}

