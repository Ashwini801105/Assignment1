package main

import (
	"fmt"
	"math/rand"
)

func main() {

	randInt := make([]int, 9)
	var i int
	randInt = append(randInt, 7)
	fmt.Println(randInt)

	randInt = append(randInt, 14)
	fmt.Println(randInt)

	randInt = append(randInt, 18)
	fmt.Println(randInt)

	for i := 0; i < 4; i++ {
		randInt[i] = rand.Intn(90)
	}

	fmt.Println(randInt[i])

}
