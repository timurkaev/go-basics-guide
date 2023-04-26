package main

import "fmt"

func main() {
	// Signed integer(Знаковые целые числа)
	var num1 int = 214748367             // целое число
	var num2 int8 = 127                  // целое число от -128 до 127
	var num3 int16 = -32768              // целое число от -32768 до 32767
	var num4 int32 = 2147483647          // целое число от -2147483648 до 2147483647
	var num5 int64 = 9223372036854775807 // целое число от -9223372036854775808  до  9 223 372 036 854 775 807
	fmt.Println("Signed integer(Знаковые целые числа)")
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println("")

	// Unsigned integer(Безнаковые целые числа)
	var num6 uint = 65535                  // безнаковое целое число
	var num7 uint8 = 255                   // от 0 до 255
	var num8 uint16 = 65535                // от 0 до 65535
	var num9 uint32 = 4294967295           // от 0 до 4 294 967 295
	var num10 uint64 = 9223372036854775807 // от 0 до 18 446 744 073 709 551 615
	fmt.Println("Unsigned integer(Безнаковые целые числа)")
	fmt.Println(num6)
	fmt.Println(num7)
	fmt.Println(num8)
	fmt.Println(num9)
	fmt.Println(num10)
	fmt.Println("")

	// Числа с плавающей точкой
	var num11 float32 = 2.10   // от 1.4*10^-45 до 3.4*10^38
	var num12 float64 = 110.12 // от 4.9*10^-324 до 1.8*10^308
	fmt.Println("Числа с плавающей точкой")
	fmt.Println(num11)
	fmt.Println(num12)
	fmt.Println("")

	// Комплексные числа
	var num13 complex64 = 4 + 2i
	var num14 complex128 = 2 + 5i
	fmt.Println("Комплексные числа")
	fmt.Println(num13)
	fmt.Println(num14)
	fmt.Println("")

	// Строки
	/*
		Строки в Go представляют собой массив из значений типа byte.
		По этой причине к элементам строки можно обращаться по индексу,
		а к самим строкам применима встроенная функция len, которая возвращает её длину в байтах
	*/
	var str string = "Hello World"
	var str2 string = `Hello
		world
	`
	var letter byte = str[0]
	var strLen int = len(str)
	fmt.Println("Строки")
	fmt.Println(str)  // Hello World
	fmt.Println(str2) /* Hello
	World */
	fmt.Println(letter) // 72
	fmt.Println(strLen) // 11
	fmt.Println("")

	// Булевые значения
	var b1 bool = false
	var b2 bool = true
	fmt.Println("Булевые значения")
	fmt.Println(b1) // false
	fmt.Println(b2) // true

}
