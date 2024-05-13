package main

import "fmt"

func main() {


	//  ||  OR
	//  &&  AND
	//  !   NOT

	val := true || false
	fmt.Println(val)

	val2 := true || false && true
	fmt.Println(val2)


// IF ELSE ,ELSE IF

	name :="tim"

	fmt.Println("before if")

 // you can change =! with == and can see different 
    if name !="tim" {
		fmt.Println("inside if")
    }
	fmt.Println("after if ")


	age := 17
	


	if age >= 18 {
		fmt.Println("you can ride !")
	} else {
		fmt.Println("you cannot ride !")
		fmt.Printf("wait %d years", 18- age )

	}

	new_age := 19

	
	if age >= 18 {
		fmt.Println("you can ride !")
	} else if  age >= 14 && new_age < 18 {
		fmt.Println("you can ride with a parent !")
		

	} else{
		fmt.Println("you cannot ride !")
		fmt.Printf("wait %d years", 18- age )
	}




}

