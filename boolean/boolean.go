package main

import (
	"fmt"
	"os"
)



func isEven(nbr int) bool {
	if even(nbr){
		return true
	} else {
		return false
	}
}

func even(nbr int) bool{
	if (nbr%2)==0{
		return true
	}
	return false
}

func main() {
	lengthOfArg:=len(os.Args[1:])
	EvenMsg:="I have an even number of arguments"
	OddMsg:="I have an odd number of arguments"
	if isEven(lengthOfArg) {
		fmt.Println(EvenMsg)
	} else {
		fmt.Println(OddMsg)
	}
}