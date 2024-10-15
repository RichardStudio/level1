/*
1. Глобальные переменные будут существовать в памяти пока вся программа не завершится, поэтому лучше объявить локально
2. Срез justString = v[:100] берет не первые 100 элементов строки, а первые 100 байтов. Нужно использовать руны
3. Из-за того, что мы используем срез, в памяти по-прежнему будет существовать огромная строка, потому что мы на нее
ссылаемся.
*/
package main

import (
	"fmt"
)

func createHugeString(size int) string {
	// Пример создания большой строки
	hugeString := make([]byte, size)
	for i := range hugeString {
		hugeString[i] = 'a'
	}
	return string(hugeString)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	return string([]rune(v)[:100]) // Создание новой строки из первых 100 символов
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}
