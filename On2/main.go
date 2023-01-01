package main

import (
	"log"
	"time"
)

type integerArray []int

/*The O(n*n) complexity is also known as quadratic complexity, it basically grows following the input length of N times N*/
func main() {
	execTime := time.Now()

	unorderedArray := integerArray{1, 3, 4, 2, 5, 6, 9, 8, 10, 11, 25, 23, 22}
	unorderedArray.nestedLoop()

	log.Printf("Executed in %v ns", time.Now().Sub(execTime).Nanoseconds())
}

/* An example of this complexity is a nested for loop, more specifically with 2 loops since it's a quadratic complexity. */
func (ar integerArray) nestedLoop() {
	arrayLenght := len(ar)
	iterationCounter := 0

	log.Println("Array length", arrayLenght)
	for i := 0; i < arrayLenght; i++ {
		for j := 0; j < arrayLenght; j++ {
			iterationCounter++
		}
	}

	log.Println("Iteration Counter:", iterationCounter)
}
