package main

import (
	"fmt"
	"strings"
)

var pword string = "pandasarecool"

func compare() {
	var pass string
	fmt.Println("Enter a password: ")
	fmt.Scan(&pass)
	var result int = strings.Compare(pass, pword)
	if result == 0 {
		fmt.Println("The password you have entered is correct.")
	} else {
		fmt.Println("The password you have entered is incorrect.")
	}
}

func main() {
	compare()
}
