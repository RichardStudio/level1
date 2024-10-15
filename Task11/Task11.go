package main

import "fmt"

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}
	res := Intersection(set1, set2)
	fmt.Println(res)
}

func Intersection(set1, set2 []int) []int {
	compute := make(map[int]bool)
	var result []int
	for _, i := range set1 {
		compute[i] = true
	}
	for _, i := range set2 {
		if compute[i] {
			result = append(result, i)
		}
	}
	return result
}
