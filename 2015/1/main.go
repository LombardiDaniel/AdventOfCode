package main

import (
	"adventOfCode-015-1/solver"
	"fmt"
	"log"
	"os"
)


func main() {
	fmt.Print("Hello, World!\n")

	inputBytes, err := os.ReadFile("./test.txt")

	if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }

	res := solver.CountSantaFloors(string(inputBytes[:]))

	fmt.Println(res)
}