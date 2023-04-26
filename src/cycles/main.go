package main

import "fmt"

func main() {
	// Бесконечный цикл
	for {
		// код, выполняемый внутри бесконечного цикла
	}

	// Трёхкомпонентный цикл
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	// Цикл с одним условием
	i := 0
	for i < 5 {
		// выводим результат на экран
		fmt.Println(i)
		// наращиваем переменную
		i++
	}

	// Цикл range
	// Цикл range используется для комплексных типов — слайса и мапы (map).
	// создаём массив
	array := [3]int{1, 2, 3}
	// итерируемся
	for arrayIndex, arrayValue := range array {
		fmt.Printf("array[%d]: %d\n", arrayIndex, arrayValue)
	}

	// Ключевые слова break и continue
	// break — выход из цикла;
	// continue — переход к следующей итерации цикла (вызов post-действия, если оно задано).
	sum, limit := 0, 100
	for i := 0; true; i++ {
		if i%2 != 0 {
			continue // переход к следующему числу, так как i — нечётное
		}

		if sum+i > limit {
			break // выход из цикла, так как сумма превысит заданный предел
		}

		sum += i
	}
	fmt.Println(sum)

	// Метки
	// В языке Go есть метки (labels), которые позволяют перемещаться к разным частям кода.
outerLoop:
	for i := 0; i < 10; i++ {
	almostInnerLoop:
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				if k == 2 {
					break outerLoop // Выходим из внешнего цикла
				}
				if k == 4 {
					break almostInnerLoop // Выходим из вложенного цикла j
				}
			}
		}
	}
}
