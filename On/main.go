package main

import (
	"log"
	"time"
)

type integerArray []int

/*The O(n) complexity is also known as linear complexity, it basically grows following the input length of N*/
func main() {
	execTime := time.Now()

	unorderedArray := integerArray{1, 3, 4, 2, 5, 6, 9, 8, 10, 11, 25, 23, 22}
	pos, ok := unorderedArray.findElement(25)
	if ok {
		log.Println("Element found at position:", pos)
	} else {
		log.Println("Element not found")
	}

	log.Printf("Executed in %v ns", time.Now().Sub(execTime).Nanoseconds())
}

/*An easy example is searching for a value in an unordered array. In that case we should look at the N element and
compare if they are the desired value. The best case scenario is when the desired value is at the first position, and the
worst at the last one. When talking about Big O notation, we always consider the worst case scenario to define an
algorithm complexity*/
func (ar integerArray) findElement(value int) (position int, exists bool) {
	for index, arrayValue := range ar {
		if arrayValue == value {
			return index, true
		}
	}

	return 0, false
}
