package main

import "fmt"

func main()  {
	// strings
	var name string = "ibrahim"
	var name1 = "ibrahim"
	var name2 string

	fmt.Println(name, name1, name2)

	name = "baba"
	name1 = "osh"
	name2 = "ade"
	fmt.Println(name, name1, name2)

	name3 := "ibro"
	fmt.Println(name3)

	// ints
	var age int = 20
	var age1 = 23
	var age2 int
	age3 := 30

	fmt.Println(age, age1, age2, age3)

	// bits & memory
	// set integer to 8 bits
	var ageRange int8 = 127
	// set integer to 16 bits
	var ageRange16 int16 = 32767
	// set integer to 32 bits
	var ageRange32 int32 = 2147483647
	// set integer to 64 bits
	var ageRange64 int64 = 9223372036854775807
	// this only take positive numbers
	var ageUint uint = 9
	// this only take positive and set up to 8 bits
	var ageUint8 uint8 = 9
	ageNormal := 40

	fmt.Println(ageRange, ageRange16, ageRange32, ageRange64, ageUint, ageUint8, ageNormal)

	// float
	var floanum float32 = 32.5
	var floanum1 float64 = 988773646536526364546565.5
	floatnormal := 2.6

	fmt.Println(floanum, floanum1, floatnormal)



}