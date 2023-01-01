package main

import (
	"log"
	"time"
)

type integerArray []int

/*The O(1) complexity describes an algorithm where the complexity is constant, that means the input size
won't impact to the function growth.*/
func main() {
	execTime := time.Now()

	arrayOfIntegers := integerArray{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arrayOfIntegers.getElementByIndex(5)

	log.Printf("Executed in %v ns", time.Now().Sub(execTime).Nanoseconds())
}

/*An easy example and the most used to describe this complexity is gathering the first element of an array,
we can simply get it using the index.*/
func (ar integerArray) getElementByIndex(index int) int {
	el := ar[index]
	log.Printf("Element at index %v is: %v\n", index, el)
	return el
}
