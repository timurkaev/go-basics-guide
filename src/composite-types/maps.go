package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["name"] = "Ismail"
	m["lastName"] = "Timurkaev"
	delete(m, "lastName")

	m2 := make(map[string]int)
	m2["price"] = 400
	m2["age"] = 23
	m2["height"] = 187
	m2["weight"] = 103

	for k, v := range m2 {
		// k будет перебирать ключи,
		// v — соответствующие этим ключам значения
		fmt.Println("key = ", k)
		fmt.Println("value = ", v)
	}

	fmt.Println(m)
	fmt.Println(m2)

}
