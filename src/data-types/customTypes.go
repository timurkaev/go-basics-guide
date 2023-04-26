package main

import "fmt"

type Name string
type LastName string

var lastName LastName
var name Name

func main() {
	lastName = "Timurkaev"
	// name = lastName  ошибка типизации
	// cannot use fruit (variable of type Fruit) as type Name in assignment

	name = Name(lastName) // так работает
	fmt.Println(lastName)
}
