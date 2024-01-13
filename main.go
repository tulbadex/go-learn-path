package main

import (
	"fmt"
	"math"
)

func greeting(n string)  {
	fmt.Printf("Good morning %v \n", n)
}

func goodBye(n string)  {
	fmt.Printf("Goodbye %v \n", n)
}

func area(r float64) float64 {
	return math.Pi * r * r
}

func callable(n []string, f func(string))  {
	for _, v := range n {
		f(v)
	}
}

func main()  {
	greeting("ibrahim")
	goodBye("ibrahim")
	callable([]string{"ibrahim", "baba", "tunde", "ade"}, greeting)
	callable([]string{"ibrahim", "baba", "tunde", "ade"}, goodBye)
	a1 := area(7.5)
	fmt.Println(a1)
	fmt.Printf("Area of a circle is %0.3f \n", a1)
}