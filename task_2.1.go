/*
Задание 1.
-Напишите программу, которая на вход получала бы строку, введённую пользователем, а в файл писала № строки, дату и сооб-
щение в формате: 2020-02-10 15:00:00 продам гараж.
-При вводе слова exit программа завершает работу.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	// var userMsg string
	count := 1

	fileName := "file.txt"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Println()
	fmt.Println("Введите сообщение с клавиатуры:")

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		userMsg := sc.Text()

		if string(userMsg) == "exit" {
			fmt.Println("Завершение работы...")
			fmt.Println()
			break
		} else {
			currentDateTime := time.Now().Format("2006-01-02 15:04:05")

			msg := fmt.Sprintf("%d %s %s\n", count, currentDateTime, userMsg)
			fmt.Println()
			fmt.Println(msg)
			_, err = io.WriteString(file, msg)
			if err != nil {
				panic(err)
			}
			count += count
			fmt.Println("Введите сообщение с клавиатуры или exit для выхода: ")
		}
	}
}
