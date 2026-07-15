package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// создание файла
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	file.Close()

	// открытие файла(только на чтение)
	file2, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	file2.Close()

	// открытие файла
	file3, err := os.OpenFile("test.txt", os.O_RDWR, 0755)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	file3.Close()

	// запись в файл
	file, err = os.OpenFile("test.txt", os.O_RDWR, 0755)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	file.Write([]byte("Dianich"))
	file.Close()

	// удобная запись строки
	file, err = os.OpenFile("test.txt", os.O_RDWR, 0755)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	file.WriteString("Dianich string")
	file.Close()

	// чтение файла целиком
	data, err := os.ReadFile("test.txt")
	fmt.Printf("From file test.txt: %s\n", string(data))

	// запись в файл
	os.WriteFile("test.txt", []byte("Write to file"), 0755)

	// буферизированное чтение
	file, err = os.Open("test2.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	// создание каталога
	err = os.Mkdir("data", 0755)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// создание вложенных каталогов
	err = os.MkdirAll(
		"data/images/users",
		0755,
	)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// удаление файла
	os.Remove("test.txt")

	// удаление каталога(перед этим создать папку)
	os.Remove("test")

	// удаление дерева каталогов
	os.RemoveAll("data")

	// переименование файла(не забыть перед этим переименовать в old)
	os.Rename("old.txt", "new.txt")

	// информация о файле
	inf, err := os.Stat("test2.txt")
	fmt.Printf("filename = %s, filesize = %d b\n", inf.Name(), inf.Size())

	// текущая директория
	dir, _ := os.Getwd()
	fmt.Printf("Current dir = %s\n", dir)
}
