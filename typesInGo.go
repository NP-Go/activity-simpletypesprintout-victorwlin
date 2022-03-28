package main

import "fmt"

//Insert variables declarations here based on activity

func tellMeTypes() {
	//insert code here to print out types of variables
	text := "The following is the account information."
	age := 20
	remainingCredits := 123.55
	subscribedUser := true

	fmt.Printf("%T %T %T %T", text, age, remainingCredits, subscribedUser)
}

func main() {
	tellMeTypes()
}
