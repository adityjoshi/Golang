package main

import "fmt"

func main() {
	age := 19
	name := "Adi"
	fmt.Println("hello")
	fmt.Println("hey i am ", name, " my age is ", age )
	
	// formated strings 
	// %v common for variable
	fmt.Printf("hey i am %v and my age is %v \n " , name, age )
	fmt.Printf("hey i am %q and my age is %v \n", name , age)
	// %q puts quotes over the string only 

	// Sprintf saves the formatted strings 
	var str = fmt.Sprintf("hey i am %v and my age is %v \n", name, age)
	fmt.Print(str)

}