package main

import "fmt"

func main()  {
	name := "ibrahim"
	age := 24
	// strings
	fmt.Print("Hello, ")
	fmt.Print("Ibrahim \n")
	fmt.Print("New Line \n")

	fmt.Println("Hello Ibrahim")
	fmt.Println("Goodbye Ibrahim")

	fmt.Println("My name is", name, "and my age is", age)

	// formatting string
	fmt.Printf("My name is %v and my age is %v \n", name, age)
	fmt.Printf("My name is %q and my age is %q \n", name, age)
	fmt.Printf("Age is type of %T \n", age)
	fmt.Printf("Your score is %f \n", 55.5)
	fmt.Printf("Your score is %0.1f \n", 55.5)

	var re = fmt.Sprintf("My name is %v and my age is %v", name, age)
	fmt.Println(re)
	fmt.Println("saved string is:", re)
}