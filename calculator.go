package main

import (
	"fmt"
)

func main() {
	var num1, num2,  result float64
	var operator string

	fmt.Println("Enter First number")
	fmt.Scanln(&num1)
	fmt.Println("Enter 2nd number")
	fmt.Scanln(&num2)
	fmt.Println("Enter operator")
	fmt.Scanln(&operator)

	switch operator {
	case "+" :
		result = num1 + num2
	case "*":
		result= num1*num2
	case "-":
		result = num1-num2
	case "/":
		if num2 !=0 {
			result = num1/num2
			
		} else {
			fmt.Println("Math Error can not divided by zero")
			return
		}
	default:
		fmt.Println("Invalid Operator")
		return
		
	}

	fmt.Println("Result of %s", operator, "is == %f", result)

}
