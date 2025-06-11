package main

import "fmt"

type Dog struct {
	name string
	age  int
}

type Speaker interface {
	Speak() string
}

// Method
func (d Dog) Speak() string {
	return "Woof!"
}

func saySomething(speaker Speaker){
	fmt.Println(speaker.Speak())
}

func main() {
	dog := Dog{"Jian",2}
	saySomething(dog)
}