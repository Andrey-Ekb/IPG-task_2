/*Задание 3. Уровни доступа
Что нужно сделать:
Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.
Рекомендация:
Для проверки создайте файл, установите режим только для чтения, закройте его, а затем, открыв, попытайтесь прочесть из
него данные.*/

package main

import (
	"fmt"
	"os"
)

func main() {

	text := "Проверка, создание файла для чтения"

	file, err := os.Create("file_test.txt")

	if err != nil {
		fmt.Println("не смогли создать файл", err)
		return
	}

	defer file.Close()
	file.WriteString(text)

	//прверяем прова нашего файла
	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Права доступа к файлу %v\n", fi.Mode())

	// Задаем прова файла только для чтения
	err = file.Chmod(0444)
	if err != nil {
		panic(err)
	}
	fi, err = file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Права доступа к файлу %v\n", fi.Mode())

}
