package main

import "fmt"

type Test struct {
	Id   int
	Name string
}

func (test *Test) ChangeTestName(newName string) {
	test.Name = newName
}

type ChangeTest interface {
	ChangeTestName(string)
}

func main() {
	// с ключевым словом type для создания нового типа
	type TestId int
	var id TestId = 2
	var x int = 5
	fmt.Printf("id = %d, x = %d\n", id, x)
	//x = id    тут компилятор ругается, так как это разные типы
	// нужно явное преобразование
	x = int(id)
	fmt.Printf("id = %d, x = %d\n", id, x)

	// теперь alias(вот это уже просто как второе имя)
	type TestId2 = int
	var id2 TestId2 = 2
	var x2 int = 5
	fmt.Printf("id2 = %d, x2 = %d\n", id2, x2)
	x2 = id2
	fmt.Printf("id2 = %d, x2 = %d\n", id2, x2)

	// struct(объявлено выше)
	test := Test{
		Id:   2,
		Name: "test2",
	}
	fmt.Printf("test = %v, type = %T\n", test, test)
	test.ChangeTestName("newTest2")
	fmt.Printf("test = %v, type = %T\n", test, test)

	// еще можно интерфейс показать(тоже выше)

	// ну и как вариант еще что-то типо перечисления
	// хотя по сути это просто набор констант
	const (
		One = iota + 1
		Two
		Three
	)
	fmt.Printf("One = %d, Two = %d, Three = %d", One, Two, Three)
}
