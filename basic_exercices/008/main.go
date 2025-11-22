// Create an array of 5 strings, and initialize it's 2 first values with some random names

package main

import "fmt"

func main () {
	// Here goes your code
	var array = [5]string{"Alice","Bob","","",""}
	array_1 := array[0]
	array_2 := array[1]
	fmt.Printf("first on : %s \nand the second one : %s",array_1,array_2)
}