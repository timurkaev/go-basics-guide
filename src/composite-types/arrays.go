package main

import "fmt"

func main() {
	var arr [7]int                // Массив фиксированной длины
	arr2 := [5]int{1, 2, 3, 4, 5} // Создание значений массива сразу
	arr3 := [7]int{1, 2, 3}
	arr4 := [6]string{"1", "2", "3"}

	// Три точки указывают компилятору, что размер массива должен быть выведен исходя из размера списка инициализации.
	arr5 := [...]string{"Go", "C++", "JavaScript"} // Создание массива по длине списка

	arr6 := [7]int{6: 11, 2: 3} // По индексу 6 будет значение 11, а по индексу 2 значение 3

	fmt.Println(arr)  // [0 0 0 0 0 0 0]
	fmt.Println(arr2) // [1 2 3 4 5]
	fmt.Println(arr3) // [1 2 3 0 0 0 0]
	fmt.Println(arr4) // [1 2 3   ]
	fmt.Println(arr5) // [Go C++ JavaScript]
	fmt.Println(arr6) // [0 0 3 0 0 0 11]

	fmt.Println(arr2[3]) // Обращение к элементам массива
	fmt.Println(arr5[0])

	// Обход массивов
	for i := 0; i < len(arr5); i++ {
		fmt.Println("for", arr5[i])
	}

	// Ключевое слово range
	for el := range arr2 {
		fmt.Println("for range", el)
	}
}
