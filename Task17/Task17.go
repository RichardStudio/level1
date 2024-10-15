package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		// Проверка, если элемент найден
		if arr[mid] == target {
			return mid
		}

		// Если целевой элемент меньше среднего, ищем в левой половине
		if arr[mid] > target {
			right = mid - 1
		} else {
			// Иначе ищем в правой половине
			left = mid + 1
		}
	}

	// Возвращаем -1 если элемент не найден
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Element %d found on pos %d\n", target, result)
	} else {
		fmt.Printf("Element %d not found\n", target)
	}
}
