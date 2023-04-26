package main

import "fmt"

/*
Операторы сравнения
	> — больше;
	< — меньше;
	>= — больше или равно;
	<= — меньше или равно;
	== — равно;
	!= — не равно.
*/

/*
А также логические операторы:
	&& — логическое И;
	|| — логическое ИЛИ;
	! — логическое НЕ.
*/

func main() {
	var a int = 1
	var b = false

	// Логическое НЕ
	if !b {
	}

	var c, d int
	// Логическое И
	if c == 1 && d == 2 {
	}
	// Логическое ИЛИ
	if c == 1 || d == 2 {
	}

	if a == 1 {
		// сценарий, если условие if выполнено
	} else if a == 2 {
		// сценарий, если условие else if выполнено
	} else {
		// сценарий, если условие else if не выполнено
	}

	// Конструкция switch case
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("Default case")
	}
}
