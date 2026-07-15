package main

import (
	"errors"
	"fmt"
)

// функция без параметров и которая ничего не возвращает
func SayHello() {
	fmt.Printf("Hello, Golang!!!\n")
}

// функция с 1 параметром
func SayHelloFor(name string) {
	fmt.Printf("Hello, %s!!!\n", name)
}

// функция с несколькими параметрами(и тут еще хотелось вот продемонстрировать короткую запись)
func SayHelloForAny(name1, name2 string) {
	fmt.Printf("Hello, %s and %s!!!\n", name1, name2)
}

// функция с возвращаемым значением
func Sum(a, b int) int {
	return a + b
}

// функция с несколькими возвращаемыми значениями
func div(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Error: division by zero")
	}
	return float64(a) / float64(b), nil
}

// функция с именованным возвращаемым значением
func Sub(a, b int) (res int) {
	res = a - b
	return
}

// функция с функцией в качестве параметра
func ExecAnotherFunc(f func()) {
	f()
}

// variadic-функция
func VariadicSum(nums ...int) (res int) {
	for _, val := range nums {
		res += val
	}
	return
}

// особая функция, точка входа в программу
func main() {
	SayHello()
	SayHelloFor("Dianich")
	SayHelloForAny("Dianich", "Dianich2")
	fmt.Printf("2 + 2 = %d\n", Sum(2, 2))
	if res, err := div(20, 0); err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("a / b = %f\n", res)
	}
	fmt.Printf("2 - 2 = %d\n", Sub(2, 2))

	// функция как значение
	f := func() {
		fmt.Printf("Func f\n")
	}

	f()

	// анонимная функция
	func(name string) {
		fmt.Printf("Hello, %s, from anonymous func\n", name)
	}("Dianich")

	// вызов функции с параметром функцией
	ExecAnotherFunc(SayHello)

	// вызов variadic-функции
	fmt.Printf("2 + 2 + 2 + 2 = %d\n", VariadicSum(2, 2, 2, 2))
}
