package main

import "fmt"

func main() {
	/*
	   Для удобного объявления и инициализации блоков констант в Go есть автоматический инкремент iota.
	   При объявлении каждого блока const значение iota равно 0 и увеличивается на 1 для каждого следующего элемента:
	*/

	const (
		white = 2*iota + 1
		black
		red
		blue
		yellow
		pink
	)

	fmt.Println(white, black, red, blue, yellow, pink)
}
