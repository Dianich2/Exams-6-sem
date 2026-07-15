package main

import "fmt"

func testFunc() bool {
	fmt.Printf("Hello from testFunc\n")
	return true
}

func main() {
	//// присваивание
	// простое присваивание
	var x int
	x = 2

	y := 22

	// множественное присваивание
	a, b := 22, 222

	// можно вот так легко свапать значения
	a, b = b, a

	//// арифметические операции
	fmt.Println("Арифметические операции")
	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("y / x = %d\n", y/x)
	fmt.Printf("x %% y = %d\n", x%y)

	//// операции сравнения
	fmt.Println("Операции сравнения")
	fmt.Printf("x > y = %t\n", x > y)
	fmt.Printf("x < y = %t\n", x < y)
	fmt.Printf("x >= y = %t\n", x >= y)
	fmt.Printf("x <= y = %t\n", y <= x)
	fmt.Printf("x == y = %t\n", x == y)
	fmt.Printf("x != y = %t\n", x != y)

	//// логические операции
	fmt.Println("Логические операции")
	fmt.Printf("x > y && x < y = %t\n", x > y && x < y)
	fmt.Printf("x > y || x < y = %t\n", x > y || x < y)
	fmt.Printf("!(x > y) = %t\n", !(x > y))

	//// короткое замыкание
	t := true || testFunc()
	fmt.Printf("t = %t\n", t)

	t = false && testFunc()
	fmt.Printf("t = %t\n", t)

	//// унарные операции
	fmt.Println("Унарные операции")
	fmt.Printf("+x = %d\n", +x)
	fmt.Printf("-x = %d\n", -x)
	fmt.Printf("&x = %p\n", &x)
	fmt.Printf("*(&x) = %d\n", *(&x))
	x++
	fmt.Printf("x++ = %d\n", x)
	x--
	fmt.Printf("x-- = %d\n", x)

	//// побитовые операции
	fmt.Println("Побитовые операции")
	fmt.Printf("x & y = %d\n", x&y)
	fmt.Printf("x | y = %d\n", x|y)
	fmt.Printf("x ^ y = %d\n", x^y)
	fmt.Printf("x << 2 = %d\n", x<<2)
	fmt.Printf("x >> 2 = %d\n", x>>2)
}
