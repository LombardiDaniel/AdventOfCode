package main

import (
	"adventOfCode-015-1/solver"
	"fmt"
	"log"
	"os"
)


func main() {
	inputBytes, err := os.ReadFile("./part1.txt")

	if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }

	res := solver.CountSantaFloors(string(inputBytes[:]))

	fmt.Printf("Part One: %d\n", res)

	inputBytes, err = os.ReadFile("./part2.txt")

	if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }

	res = solver.NavigateSantaFloors(string(inputBytes[:]))

	fmt.Printf("Part Two: %d\n", res)

}