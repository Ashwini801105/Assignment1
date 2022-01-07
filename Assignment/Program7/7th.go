package main

import "fmt"

func main() {

	var key string
	var mp map[string]int = map[string]int{
		"rock": 22,
		"sand": 7,
		"mud":  15,
	}
	fmt.Println(mp["sand"])
	mp["slate"] = 23
	delete(mp, "mud")

	fmt.Println(mp)
	for key = range mp {
		delete(mp, key)
	}

	fmt.Println(mp)

}
