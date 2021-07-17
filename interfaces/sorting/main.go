package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s StringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	names := StringSlice{"Tom", "Scott", "rand"}
	fmt.Printf("Before: \t%v:%T\n", names, names)
	sort.Sort(StringSlice(names))
	fmt.Printf("After:  \t%v:%T\n", names, names)
	newNames := []string{"Tom", "Scott", "randy"}
	fmt.Println("Standard library for sort also provides native type sorting")
	fmt.Printf("Before: \t%v:%T\n", newNames, newNames)
	sort.Strings(newNames)
	fmt.Printf("After:  \t%v:%T\n", newNames, newNames)
}
