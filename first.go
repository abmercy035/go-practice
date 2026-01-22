package main

import "fmt"

func main() {

	x, y, name := 14, 15, "Abraham Samuel"
	const pi float64 = 3.16412715

	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x % y)

	isbool := true
	hate := false

	fmt.Println(isbool && hate)
	fmt.Println(isbool || hate)
	fmt.Println(!isbool)

	fmt.Println(len(name))

	// const pi float64 = 3141512435
	// var name string = "Abraham Samuel"
	// fmt.Println(len(name))
	// fmt.Printf("%.4f \n", pi)

	// Pointers
	for i := 1; i < 5; i++ {
		changeX(&x, i)
	}
	fmt.Println("st", x)

	//Printf - formatted printing

	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", name)
	fmt.Printf("%t \n", hate)
	fmt.Printf("%d \n", x)
	fmt.Printf("%b \n", 34)
	fmt.Printf("%c \n", 36)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%e \n", pi)

	// Loops

	for i := 1; i < 8; i++ {
		for j := 1; j < i; j++ {

			fmt.Printf(" * ")
		}
		fmt.Println()
	}

	b := 1

	for b <= 6 {
		fmt.Println(b)
		b++
	}
	// Decision Making
	age := 8

	if age > 8 {
		fmt.Println("Yes you can play")
	} 	else {
		fmt.Println("sorry you cant play")
	}

	switch age {
		case 8: fmt.Println("You go and read")
		case 18: fmt.Println("You go to college")
		default: fmt.Println("You free to go")
	}

//Array 

var numArry[10] int

numArry[0] = 23
numArry[1] = 45
numArry[2] = 67
numArry[3] = 89
fmt.Println(numArry[2])

someArray := [2]int{1,2}

fmt.Println(someArray)

for i, val := range someArray{
	fmt.Println(val, i)
}
// Array slice

toSlice := []int{3,5,6,7, 9,8}

sliced := toSlice[1:4]
	fmt.Println(sliced)

newSlice:= make([]int, 5, 10)

copy(newSlice, toSlice)

rdSlice := append(newSlice, 45, 67, 89, -1)

	fmt.Println(newSlice)
	fmt.Println(rdSlice)

	fmt.Println(toSlice)

}

func changeX(x *int, num int) {
	if num == 2 {
		*x = num
	}
}
