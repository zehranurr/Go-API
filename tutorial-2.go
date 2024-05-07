package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Printf("hello  %T  %v", 10, 20)

	// %v value in default format

	// %T type

	// %% literal %

	// %t true or false

	// %b : base 2

	// %o : base 8

	// %d : base 10

	// %x : base 16

	// %e : scientific notation

	// %f :  decimal no exponent

	// %x : base 16

	// %s string

	//user input

	/* scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type somethng ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf(" you typed : %q", input) */

	//if you wanna convert input type use strconv
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Printf("Type ,the year you were bron : ")
	scanner.Scan()
	input2, _ := strconv.ParseInt(scanner.Text(),10,64)
	fmt.Printf(" you will be %d years old at the end of 2000", 2020- input2)






}
