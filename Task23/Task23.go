package main

import "fmt"

func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	i := 2 // индекс элемента для удаления
	fmt.Println("input:", slice)
	slice = remove(slice, i)
	fmt.Println("After removing:", slice)
}
