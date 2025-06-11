package main

import "fmt"

//* these structs can only be used in this file

//* for external use make sure all fields are starting with capital letter

type Subject struct {
	title         string
	firstGrading  float64
	secondGrading float64
	thirdGrading  float64
	fourthGrading float64
}

type Person struct {
	name     string
	age      int
	year     string
	birthday string
	subjects []Subject
}

//* *Person - passing a pointer avoids copying the whole Person struct - which saves memory and time
func ChangeName(person *Person, newName string) {
	person.name = newName
}

func main() {
	//* slice of subjects
	subjects := []Subject{
		{
			title:         "Math",
			firstGrading:  75,
			secondGrading: 85,
			thirdGrading:  96,
			fourthGrading: 70,
		},
		{
						title:         "English",
			firstGrading:  75,
			secondGrading: 80,
			thirdGrading:  85,
			fourthGrading: 90,
		},
	}
	person := Person{name: "Justine", age: 25, year: "first year", birthday: "May 12,1990", subjects: subjects}

	fmt.Println(person)

	//* &person pass address of x
	ChangeName(&person, "Justine Hornales")

	fmt.Println("New name:", person)
}
