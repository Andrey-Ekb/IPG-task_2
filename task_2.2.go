/*Задание 2. Интерфейс io.Reader
Что нужно сделать:
-Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике, без использо-
вания ioutil. Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.
Рекомендация:
Для получения размера файла воспользуйтесь методом Stat(), который возвращает информацию о файле и ошибку.*/

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return
	}

	defer file.Close()

	buf := make([]byte, 256)
	_, err = file.Read(buf)
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return
	}
	fmt.Println(string(buf))
}
