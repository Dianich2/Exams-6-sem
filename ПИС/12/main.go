package main

import "fmt"

type Test struct {
	Id   int
	Name string
}

func (test Test) ChangeTestNameByVal(newName string) {
	test.Name = newName
}

func (test *Test) ChangeTestNameByPtr(newName string) {
	test.Name = newName
}

// это для демонстрации встраивания
type Person struct {
	Name string
}
type Employee struct {
	Person
	Salary int
}

func main() {
	// объявление базовое(каждое поле получит Zero Value значение)
	var test1 Test
	fmt.Printf("test1 = %v, type = %T\n", test1, test1)

	// объявление с заданием параметров по именам
	// можно в любом порядке
	test2 := Test{
		Name: "test2",
		Id:   2,
	}
	fmt.Printf("test2 = %v, type = %T\n", test2, test2)

	// объявление с заданием параметров по порядку
	// тогда нужно как при создании структуры
	test3 := Test{
		3,
		"test3",
	}
	fmt.Printf("test3 = %v, type = %T\n", test3, test3)

	// сравнение структур(только если все поля сравнимы)
	fmt.Printf("test2 == test3 = %t\n", test2 == test3)

	// метод по значению(смотреть выше)
	fmt.Printf("test2 = %v, type = %T\n", test2, test2)
	test2.ChangeTestNameByVal("test22")
	fmt.Printf("test2 = %v, type = %T\n", test2, test2)
	test2.ChangeTestNameByPtr("test22")
	fmt.Printf("test2 = %v, type = %T\n", test2, test2)

	// встраивание структур
	emp := Employee{
		Person: Person{
			Name: "Dianich",
		},
		Salary: 2222,
	}

	fmt.Printf("emp = %v, type = %T\n", emp, emp)
	// покажем, что можно именно при встраивании обращаться напрямую к полям вложенной структуры
	fmt.Printf("emp.Name = %s\n", emp.Name)
}
