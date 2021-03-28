package main

import (
	"fmt"
	"strconv"
)

func calculateX(numb1, op1, numb2, op2, numb3 string) string{
	number1, err := strconv.ParseInt(numb1, 10, 64)
	if err != nil {
		return "result = validation error"
	}

	number2, err := strconv.ParseInt(numb2, 10, 64)
	if err != nil {
		return "result = validation error"
	}

	number3, err := strconv.ParseInt(numb3, 10, 64)
	if err != nil {
		return "result = validation error"
	}

	var result int64
	var x string

	if op1 == "+" && op2 == "+" {
		result = number1 + number2 + number3
	} else if op1 == "+" && op2 == "-" {
		result = number1 + number2 - number3
	} else if op1 == "-" && op2 == "+" {
		result = number1 - number2 + number3
	} else if op1 == "-" && op2 == "-" {
		result = number1 - number2 - number3
	} else {
		return "result = validation error"
	}
	
	x = fmt.Sprintf("result = %s ", fmt.Sprint(result))
	return x

}

func main() {
	fmt.Println(calculateX("5", "+", "2", "-", "1")) 
	fmt.Println(calculateX("adas5", "+", "2", "-", "1")) 
}