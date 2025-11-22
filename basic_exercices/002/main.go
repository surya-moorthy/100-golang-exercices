// Exercise: Create a program that shows your name, address and also has some comments in it
// - Use the fmt package

package main

import "fmt"

func main () {
	// Here goes your code
    name := "Surya"
	address := "No. 561, 2nd Street, Saravana Nagar,Anaicut Road, 632513"
	comments := "I am passionate developer"

	fmt.Printf(" name : %s \n address : %s \n comments : %s \n",name,address,comments)
}