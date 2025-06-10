package main

import "fmt"


func Power(number float64, power int ) float64 {
	product := number

	for ctr:=1; ctr < power; ctr++ {
		product *= number
	}

	return product

}


func main(){
	fmt.Println("Result: ", Power(4, 2))
}