package main

import "fmt"

func main() {

	age := "16"
	name := "Abraham"
	
		//# Print

	fmt.Print("hello \n")
	fmt.Print("world")

		//# Println

	fmt.Println("hello world")
	fmt.Println("a new line automatically")
	fmt.Println("my age is", age, "my name is", name)

		//# Printf (formatted strings)
	// %_ = format specifier

	fmt.Printf("my age is %v and my name is %v \n", age, name) // format arg as variables
	fmt.Printf("my age is %q and my name is %q \n", age, name) // format arg as double quoted string
	fmt.Printf("my age is %T and my name is %T \n", age, name) //	format arg as type of variable
	fmt.Printf("my age is %f \n", 4.566)                       //	format arg as float
	fmt.Printf("my age is %0.2f \n", 4.566)                    //	format arg as float into 2 decimal place

		//# Sprintf	(string format)

	var str = fmt.Sprintf("my age is %q and my name is %q \n", age, name) // format arg as double quoted string

	fmt.Print("value of str is: ", str)

	//# strings

	// var nameOne string  = "45"
	// var nameTwo = "auto string"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne ="peach"
	// nameThree = "broswer"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "fourthString"

	// fmt.Println(nameFour)

	//# ints

	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	//# bits and memory

	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint16 = 256

	// fmt.Println(numOne, numTwo, numThree)

	//# float

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 253465772345672243233333333.7
	// scoreThree := 1.5

	// fmt.Println(scoreOne, scoreTwo, scoreThree)

}
