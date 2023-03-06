package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter your age")
	var age int
	fmt.Scan(&age)
	//fmt.Println(age)

	if age < 18 {
		fmt.Println("You are not 18 years old")
	} else {
		fmt.Println("You are 18 years old")
	}
}