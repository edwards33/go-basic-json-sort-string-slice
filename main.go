package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"William", "Jack", "Liam", "Charlie", "Harry"}
	fmt.Println(str)
	sort.Sort(sort.StringSlice(str))
	fmt.Println(str)

}
