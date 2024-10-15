package main

import (
	"fmt"
)

func setBit(num int64, pos uint, value bool) int64 {
	if value {
		// Устанавливаем бит в 1
		num |= 1 << pos
	} else {
		// Устанавливаем бит в 0
		mask := ^(1 << pos)
		num &= int64(mask)
	}
	return num
}

func main() {
	var num int64  // исходное число
	var pos uint   // позиция бита, который хотим изменить
	var value bool // значение (true для 1, false для 0)
	var scanVal int

	fmt.Println("Enter a number: ")
	fmt.Scan(&num)

	fmt.Println("Enter a position: ")
	fmt.Scan(&pos)

	fmt.Println("Enter 1 to set bit to 1, 0 otherwise: ")
	fmt.Scan(&scanVal)
	// Без проверок для удобства
	if scanVal == 1 {
		value = true
	} else {
		value = false
	}

	fmt.Printf("Исходное число: %064b\n", num)
	num = setBit(num, pos, value)
	fmt.Printf("После установки %d-го бита в %t: %064b\n", pos, value, num)
}
