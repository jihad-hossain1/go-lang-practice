package main

import "fmt"

/*--------  Normal function  -------*/ 
func console (val int){
	fmt.Println(val)
}

/*-----  Return single value  -------*/
func number (num1 int, num2 int) int {
	sum := num1 + num2

	return  sum
}

/*-----  Return Multiple value  -------*/
func numbers (num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2

	return  sum, mul
}


func main () {
	month := 7
	year := 10
	
	fmt.Println(numbers(month, year))

	p, q := numbers(month, year)
	fmt.Println(p,q)
}

// this function every time execute first, before main function run
func init(){
	fmt.Println("I am execute first!")
}