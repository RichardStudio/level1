package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	res := MakeSet(data)
	fmt.Println(res)
}

func MakeSet(arr []string) []string {
	set := make(map[string]struct{})
	var result []string
	for _, s := range arr {
		set[s] = struct{}{}
	}
	for s := range set {
		result = append(result, s)
	}
	return result
}
