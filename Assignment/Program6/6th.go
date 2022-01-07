package main

import "fmt"

func main() {
	var largerNumber, temp int
	array := []int{5, 25, 30, 40, 51}

	for i, element := range array {
		if element > i {
			temp = element
			largerNumber = temp
		}
	}
	fmt.Println("Largest number of Array/slice is  ", largerNumber)
}
