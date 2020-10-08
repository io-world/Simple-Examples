package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]
	fmt.Println("We got what you need", argsWithoutProg[0])
	fmt.Println("We got what you need", argsWithoutProg[1])
}
