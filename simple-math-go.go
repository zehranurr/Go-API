package main

import "fmt"

func main() {
	var num1 int = 8
	var num2 int = 4
	// simple operator  + , - ,/, *
	answer := num1 + num2
	answer1 := num1 / num2

	fmt.Println(answer) 
	fmt.Println(answer1) 

	x := 5
    y := 6


	val :=  x < 5
	val2 := y > 3

	z := 4
	k :=  6.5
	dnm := float64(z) != float64(k)

	 // less <= greater
	 // greater >= less 


	fmt.Println(val)
	fmt.Println(val2)
	fmt.Println(dnm)
}
