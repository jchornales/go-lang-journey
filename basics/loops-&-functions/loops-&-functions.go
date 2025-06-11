package main

import "fmt"


func Power(number float64, power int ) float64 {
	product := number

	for ctr:=1; ctr < power; ctr++ {
		product *= number
	}

	return product

}

func Add(a,b float64) float64{
	return a+b
}


func Subtract(a,b float64) float64{
	return a-b
}



func Multiply(a,b float64) float64{
	return a*b
}



func Divide(a,b float64) float64{
	return a/b
}


func Calculate(a float64,b float64, operation string)  {
	var result float64

	switch operation {
	case "+":
		result= Add(a,b)
	case "-":
		result= Subtract(a,b)
	case "*":
		result= Multiply(a,b)
	case "/":
		result=Divide(a,b)
	default:
		fmt.Println("Please select operator")
	}
	// No changes needed here for rune literals.
	fmt.Println("Result", result)
}


func main(){
	Calculate(5,2, "/" )
}