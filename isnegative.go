package piscine

import "fmt"



func IsNegative(nb int) {

	if nb < 0 {
		fmt.Print("T\n")
	} else {
		fmt.Print("F\n")
	}
}