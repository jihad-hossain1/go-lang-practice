package main

import "fmt"

func main() {
	// basic variable declare
	// basicSyntax()
	// forLoop()

	slice := []string{"apple", "banana", "orange"}
	for index, value := range slice {
		fmt.Printf("index: %d, value: %s\n", index, value)
	}
}

func forLoop() {
	i := 0
	for i < 5 {
		fmt.Println("value:", i)
		i++
		if i == 3 {
			break
		}
	}
}

func basicSyntax() {
	var x int = 10
	var y string = "Variable"
	var z bool = true
	// var abc string = "local"
	const localVar string = "local const"

	fmt.Println(x, y, z)
	fmt.Println(localVar)

	// type infer variable declare
	a := 10
	b := 20
	name := "Go"
	isCool := true

	person_name := "jihad"
	age := 32
	city := "Dhaka"

	fmt.Println(person_name, age, city)
	fmt.Println(a, b, name, isCool)
}
