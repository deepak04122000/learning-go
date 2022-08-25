package functions

import "fmt"

//  Creating Addition Function
func Addition(x int, y int )int{
	return x + y
}


// Creating Function Using Loop
func PrintTable(x int) {
	for i :=1; i <= 10; i++ {
		fmt.Println(x*i)
	}
}

// Creating Function for Checking greater than
func IsGreater(x int, y int) bool {
	if x > y {
		return x > y
	}

	return false
}