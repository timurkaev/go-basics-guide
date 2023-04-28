package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Получаем читателя пользовательского ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	cnt := 0
	for {
		fmt.Print("-> ")
		// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&cnt) // Создаем указатель

		fmt.Printf("User input %d lines\n", cnt)
	}
}

func f(n *int) { // Получаем значение на которое указывает переменна &cnt выше
	*n++
}
