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

}

func changeX(x *int, num int) {
	if num == 2 {
		*x = num
	}
}
