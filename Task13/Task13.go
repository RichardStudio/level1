package main

import "fmt"

func main() {
	a := 5
	b := 10
	fmt.Printf("a: %d, b: %d\n", a, b)
	swap3(&a, &b)
	fmt.Printf("a: %d, b: %d\n", a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b *int) {
	*a += *b
	*b = *a - *b
	*a = *a - *b
}

func swap3(a, b *int) {
	*a ^= *b
	*b = *a ^ *b
	*a ^= *b
}
