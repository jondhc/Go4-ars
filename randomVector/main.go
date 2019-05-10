package main

import (
	"fmt"
	"math/rand"
)

func generateVector(size int) {
	generatedVector := make([]int, size) //slice for dinamic size array, allocates 0s
	for i := 0; i < size; i++ {
		generatedVector[i] = rand.Intn(100)
	} //end for
	fmt.Printf("%v", generatedVector)
} //end generateVector

func main() {
	generateVector(4)
} //end main
