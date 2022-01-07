package main

import "fmt"

func main() {
	var Month string = "Feb"
	switch {
	case Month == "Jan":
		fmt.Println("January")
		fallthrough
	case Month == "Feb":
		fmt.Println("February")
		fallthrough
	case Month == "Mar":
		fmt.Println("March")
		fallthrough
	case Month == "Apr":
		fmt.Println("April")
		fallthrough
	case Month == "May":
		fmt.Println("May")

	}
}
