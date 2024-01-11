package main

import (
	"fmt"
)

func main()  {
	x := 0
	for i :=0; i < 6; i++ {
		fmt.Println("value of i is", i)
	}

	for x < 6 {
		fmt.Println("value of x is", x)
		x++
	}

	names := []string{"ibrahim", "bobo", "ibro", "tunde", "baba"}
	for i :=0; i < len(names); i++ {
		fmt.Println("value of name is", names[i])
	}

	for index, value := range names {
		fmt.Printf("index is %v and value is %v \n", index, value)
	}

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
	}
}