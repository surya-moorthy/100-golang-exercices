// Exercise: With a single string variable, use access the first character with the string index
// This is only valid with ASCII characters, and the print value will be the ASCII number
// string[n] is how you should access the value

package main

import "fmt"

func main () {
	// Here goes your code
	var string1 string
	string1 = "helloWorld"

	for idx,value := range string1 {
		fmt.Println(idx,value)
	}

	fmt.Printf("the first index string ASCII value %d",string1[0])
	// ...
}
