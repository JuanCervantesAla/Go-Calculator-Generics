package main

import (
	"calculator/models"
	"fmt"
)

func main() {
	calculator := models.Calculator[int]{5, 10}
	var option int

	for {
		fmt.Print("Option Menu\n")
		fmt.Print("1-Adding\n")
		fmt.Print("2-Substracting\n")
		fmt.Print("3-Multiply\n")
		fmt.Print("4-Divide\n")
		fmt.Print("8-End\n")
		fmt.Scan(&option)

		switch option {
		case 1:
			SetNumbers(&calculator)
			var response int = calculator.Adding()
			fmt.Println(response)
		case 2:
			SetNumbers(&calculator)
			var response int = calculator.Substracting()
			fmt.Println(response)
		case 3:
			SetNumbers(&calculator)
			var response int = calculator.Multiply()
			fmt.Println(response)
		case 4:
			SetNumbers(&calculator)
			var response int = calculator.Divide()
			fmt.Println(response)
		case 8:
			return
		default:
			fmt.Println("Wrong option")
		}

		fmt.Print("\n\n\n\n")

	}
}

func SetNumbers(calculator *models.Calculator[int]) {
	fmt.Println("Enter a number: ")
	fmt.Scan(&calculator.A)
	fmt.Println("Enter another One: ")
	fmt.Scan(&calculator.B)
}
