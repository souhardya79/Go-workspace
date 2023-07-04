package main

import "fmt"

func main() {
	var a string
	var b string
	fmt.Print("Ente the first name ")
	fmt.Scan(&a)
	fmt.Println("Your full name is: " + a)
	fmt.Print("Enter the second name: ")
	fmt.Scan(&b)
	fmt.Println("The second name is : " + b)

}
