package main

import "fmt"

func main()  {
	var ages [3]int = [3]int{10, 20, 30}
	var ages1 = [4]int{10, 20, 30, 40}

	names := [4]string{"ibrahim", "bobo", "ade", "baba"}

	fmt.Println(ages, len(ages))
	fmt.Println(ages1, len(ages1))
	fmt.Println(names, len(names))

	// slices
	var scores = []int{100, 50, 60}
	scores[2] = 24
	scores = append(scores, 85, 30)
	fmt.Println(scores, len(scores))

	// slice and ranges
	ranges := names[:2]
	fmt.Println(ranges)
}