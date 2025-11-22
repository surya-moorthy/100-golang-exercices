// Exercise: User input
// Using only the fmt package, ask a user for it's name and then for it's surname
// Store it in 2 variables called "name" and "surname"
// After user has entered the data, print it out

// Tip: https://pkg.go.dev/fmt#hdra-Scanning

package main

import "fmt"

func main () {
	// Here goes your code
	var name string
	var surname string
	fmt.Println("enter your name and surname:")
	fmt.Scanf("%s %s",&name,&surname)
    fmt.Printf("Your name %s and your surname %s",name,surname)
}
