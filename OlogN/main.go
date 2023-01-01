package main

import (
	"log"
	"time"
)

type integerArraySorted []int

/*The O(log N) is the best case scenario for the real world algorithms. It would be amazing if all our algorithms were
constants O(1) but that doesn't happen really often. Most of our tasks will require more work than simply getting
a value from a map/array.*/
func main() {
	execTime := time.Now()

	sortedArray := integerArraySorted{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110}
	elementPosition, exists := sortedArray.binarySearch(0)

	if exists {
		log.Println("Element Position:", elementPosition)
	} else {
		log.Println("Element not found")
	}

	log.Printf("Executed in %v ns", time.Now().Sub(execTime).Nanoseconds())
}

/*This complexity will growth in a log n expression where N is the input size. An example of this complexity is searching
for a number in an ordered array using Binary Search. We don't have to go through the N array elements, we start from
the middle and check if our number is at left or right.*/
func (ar integerArraySorted) binarySearch(element int) (elementPosition int, elementExists bool) {
	lowerIndex := 0
	higherIndex := len(ar) - 1

	numOperations := 0
	for lowerIndex <= higherIndex {
		middleIndex := (higherIndex + lowerIndex) / 2

		switch {
		case ar.isEqual(middleIndex, element):
			return middleIndex, true
		case ar.isAtLeft(middleIndex, element):
			higherIndex = middleIndex - 1
		case ar.isAtRight(middleIndex, element):
			lowerIndex = middleIndex + 1
		}

		numOperations++
		log.Println("Number of operations:", numOperations)
	}

	return 0, false
}

func (ar integerArraySorted) isAtLeft(middleIndex, element int) bool {
	return element < ar[middleIndex]
}

func (ar integerArraySorted) isAtRight(middleIndex, element int) bool {
	return element > ar[middleIndex]
}

func (ar integerArraySorted) isEqual(middleIndex, element int) bool {
	return element == ar[middleIndex]
}
