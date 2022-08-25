package main

import (
	"fmt"
	"time"

	// importing functions folder as a package
	"github.com/deepak04122000/learning-go/functions"
)

func main() {
	// Printing Helloe world
	fmt.Print("Hello World!")

	// Defining variable with value and data types
	name := "Ram"
	fmt.Println(name)

	// Changing variable value and printing in new line
	name = "Syam"
	fmt.Println(name)

	// Declaring constant variable
	const age = 15
	fmt.Println(age)
	  
	// using time module
	fmt.Println(time.Now())
	fmt.Println(time.Hour)


	// Executing functions of functions folder
	fmt.Println(functions.Addition(2, 6))

	functions.PrintTable(10)

	fmt.Println(functions.IsGreater(9,9))
}
