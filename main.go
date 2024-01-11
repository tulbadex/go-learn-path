package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {
	greeting := "Good day, Ibrahim"
	fmt.Println(strings.Contains(greeting, "Good"))
	fmt.Println(strings.ReplaceAll(greeting, "Good", "Do you have a good"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "oo"))

	ages := []int{24, 34, 56, 12, 3, 45, 67, 88, 98}
	sort.Ints(ages)
	fmt.Println((ages))

	index := sort.SearchInts(ages, 24)
	fmt.Println((index))

	names := []string{"ibrahim", "ade", "bobo", "osho", "baba"}
	sort.Strings(names)
	fmt.Println((names))

	fmt.Println(sort.SearchStrings(names, "ibrahim"))
}