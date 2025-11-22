// Exercise: User input
// Get a number from the console and check if it's odd
// You can create another function or do it everything in the main func :) 

package main

import "fmt"

func main () {
	// Here goes your code
	var num int
    fmt.Scanf("%d",num)
	if num%2 != 0 {
		fmt.Print("it is odd")
	}else {
		fmt.Print("it is even")
	}
}