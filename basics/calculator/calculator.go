package main

import (
	"flag"
	"fmt"
)

type Values struct {
	FirstNumber  float64
	SecondNumber float64
	Operation string
}

func (values Values) Add() float64 {
	return values.FirstNumber + values.SecondNumber
}

func (values Values) Subtract() float64 {
	return values.FirstNumber - values.SecondNumber
}

func (values Values) Multiply() float64 {
	return values.FirstNumber * values.SecondNumber
}

func (values Values) Divide() float64 {
	// This block of codes is suggested by AI to make it better at handling error
	if values.SecondNumber == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	//
	return values.FirstNumber / values.SecondNumber
}

func main(){
	firstNumber := flag.Float64("firstNumber", 0, "Third Number: " )
	secondNumber := flag.Float64("secondNumber", 0,"Second Number: " )
	operator := flag.String("operator", "", "Operation:" )
	flag.Parse()

	values := Values{*firstNumber, *secondNumber, *operator}

	switch(*operator){
	case "+":
		fmt.Println(values.Add())
	case "-":
		fmt.Println(values.Subtract())
	case "*":
		fmt.Println(values.Multiply())
	case "/":
		fmt.Println(values.Divide())
		// this block of code as well
	default:
	fmt.Println("Error: Unknown operator. Use +, -, *, or /")
	}
}