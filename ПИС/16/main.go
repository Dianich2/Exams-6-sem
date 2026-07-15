package main

import "fmt"

func main() {
	// классический for
	fmt.Println("for")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	// for как while
	fmt.Println("for как while")
	x := 2
	for x < 10 {
		fmt.Printf("%d\n", x)
		x++
	}

	// бесконечный цикл
	// for{}

	// операторы управления циклом
	fmt.Println("break и continue")
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Printf("%d\n", i)
	}

	// вложеннные циклы
	fmt.Println("Вложенные циклы")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}

	// for ... range
	s := []int{2, 22, 222}
	for idx, val := range s {
		fmt.Printf("s[%d] = %d\n", idx, val)
	}

	// метки(часто используются, чтобы сразу выйти из нескольких циклов)
	fmt.Println("Метки")
Outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == 1 && j == 1 {
				break Outer
			}
		}
	}

}
