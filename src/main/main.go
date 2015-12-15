package main

import (
	"fmt"
)

func main() {

	fmt.Println("play(Y/N)?")
	var input string
	fmt.Scanln(&input)
	determine_output(&input)
}

func determine_output(answer *string) bool {
	switch *answer {
	case "Y":
		return true
	case "N":
		return false
	default:
		fmt.Println("Please Input Y/N")
		fmt.Scanln(answer)
		return determine_output(answer)
	}
}
