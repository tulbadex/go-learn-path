package main

import (
	"fmt"
)

func main()  {
	age := 33

	fmt.Println(age <= 40)
	fmt.Println(age >= 40)
	fmt.Println(age == 33)
	fmt.Println(age != 40)

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 32 {
		fmt.Println("Age is less than 32")
	} else {
		fmt.Println("Age is not less than 33")
	}

	names := []string{"ibrahim", "bobo", "baba", "tunde"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continue at index pos", index)
			continue
		}

		if index == 3 {
			fmt.Println("break at index pos", index)
			break
		}
		fmt.Printf("The value at pos %v is %v \n", index, value)

	}
	
}