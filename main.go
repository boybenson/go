package main

import (
	"fmt"
	"strconv"
)

func calculateTotalPrice(initialPrice float64, itemPrice float64) (total float64, word string) {
	total= initialPrice + itemPrice
	word = "The total Price is " + strconv.FormatFloat(total, 'f', 2, 64)
	return 
}

type Person struct {
	name string
	age int
	job string
	salary int 
}


func main (){

	// This right here is 
	// a multiline comment in Golang
	// which will be ignored.

	myname := "Benson"
	age := 24
	graduated := true
	var signed uint64 = 5000


	const appName string = "Million Dollar Project"

	fmt.Printf("My name is %v and i am %v years of age and it is %v that i am a graduate working on %v", myname, age, graduated, appName)
	fmt.Println(signed)

	var names = [3]string{"Peter", "Paul", "Jesus"}
	ages := [3]uint{10, 20, 30}

	fmt.Println(names)
	fmt.Println(ages)
	fmt.Println(len(names))

	slice1 := []int{1,2,3}
	slice2 := []int{4,5,6}

	slice3 := append(slice1, slice2...)

	fmt.Println(slice3)

	temperature:= 2

	if temperature > 20{
		fmt.Println("The weather is a little warm and better today")
	}else {
		fmt.Println("The weather is very cold ")
	}
	
	for i:= 0; i < 5; i++ {
		if i == 2{
			continue
		}
		fmt.Print(i)
	}

	for index, value := range names{
		fmt.Printf("Item with index %v has a value of %v \n", index, value)
	}

	total, word:= calculateTotalPrice(100, 200)
	fmt.Print(total)
	fmt.Print(word)

	


}
