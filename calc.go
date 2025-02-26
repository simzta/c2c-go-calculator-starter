package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func multi(a int, b int) int {
	return a * b
}

//1- create a division function (div) with 2 integer parameters and a float return

func goCal(firstNum int, secondNum int) {
	fmt.Println(firstNum, "+", secondNum, "=", plus(firstNum, secondNum))
	fmt.Println(firstNum, "-", secondNum, "=", minus(firstNum, secondNum))
	fmt.Println(firstNum, "*", secondNum, "=", multi(firstNum, secondNum))

	if secondNum == 0 {
		fmt.Println("not divisible by zero.")
	} else {
		fmt.Println(firstNum, "/", secondNum, "=", div(firstNum, secondNum))
	}
}

func main() {
	var firstNum int
	var secondNum int
	var name string

	fmt.Print("\nHello, What's your name? ")
	fmt.Scanf("%s", &name)
	fmt.Printf("\nWelcome, %s! Let's do some basic calculations.\n\n", name)

	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d", &firstNum)

	//2- ask the user to enter the second number and store the value in "secondNum"

	//3- Call the "GoCal" function with the proper parameters
}
