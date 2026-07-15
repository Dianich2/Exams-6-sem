package main

import "fmt"

type Test struct {
	Id   int32
	Name string
}

func (test2 *Test) ChangeTestName(newName string) {
	test2.Name = newName
}

func main() {
	var x = 2
	var p = &x
	fmt.Printf("Адрес = %p, значение = %d\n", p, *p)

	// тут будет panic, потому что указатель nil
	var p2 *int
	fmt.Printf("Адрес = %p\n", p2)
	//fmt.Printf("Значение = %d\n", *p2)

	// можно создавать вот так указатель и тогда он не nil
	p3 := new(int)
	fmt.Printf("Адрес = %p, значение = %d, тип = %T\n", p3, *p3, p3)

	// указатели на структуры (структура выше)
	test := Test{
		Id:   1,
		Name: "Test1",
	}

	pTest := &test
	fmt.Printf("Адрес = %p, значение = %+v, тип = %T\n", pTest, *pTest, pTest)
	fmt.Printf("Id = %d, Name = %s\n", pTest.Id, pTest.Name) // разыменование автоматически

	// покажем еще метод структуры с указателем (метод выше)
	test.ChangeTestName("Test2")
	fmt.Printf("Адрес = %p, значение = %+v, тип = %T\n", pTest, *pTest, pTest)
	fmt.Printf("Id = %d, Name = %s\n", pTest.Id, pTest.Name)
}
