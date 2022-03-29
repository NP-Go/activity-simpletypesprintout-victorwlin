package main

import "fmt"

//Insert variables declarations here based on activity
const text string = "The following is the account information."
const firstName string = "Luke"
const lastName string = "Skywalker"
const age uint = 20
const weight float32 = 73.0
const height float32 = 1.72
const remainingCredits float32 = 123.55
const accountName string = "admin"
const accountPassword string = "password"
const subscribedUser bool = true

func tellMeTypes() {
	//insert code here to print out types of variables
	fmt.Printf("type: %T  value: %v\n", text, text)
	fmt.Printf("type: %T  value: %v\n", firstName, firstName)
	fmt.Printf("type: %T  value: %v\n", lastName, lastName)
	fmt.Printf("type: %T  value: %v\n", age, age)
	fmt.Printf("type: %T  value: %v\n", weight, weight)
	fmt.Printf("type: %T  value: %v\n", height, height)
	fmt.Printf("type: %T  value: %v\n", remainingCredits, remainingCredits)
	fmt.Printf("type: %T  value: %v\n", accountName, accountName)
	fmt.Printf("type: %T  value: %v\n", accountPassword, accountPassword)
	fmt.Printf("type: %T  value: %v\n", subscribedUser, subscribedUser)
}

func main() {
	tellMeTypes()
}
