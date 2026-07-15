package main

import "fmt"

func main() {
	// ну по факту тут if-else(else if) по стандарту
	num := 19
	if num <= 10 {
		fmt.Printf("num <= 10\n")
	} else if num <= 20 {
		fmt.Printf("num <= 20\n")
	} else {
		fmt.Printf("num > 20\n")
	}

	// можно объявить переменную в самом if и она будет доступна только там
	if x := 10; x > 5 {
		fmt.Printf("x in if > 5\n")
	}

	// конструкция switch(покажем сразу все возможности)
	num2 := 6
	switch num2 {
	case 1, 2, 3, 4, 5:
		fmt.Printf("num2 <= 5\n")
	case 6:
		fmt.Printf("num2 = 6\n")
		fallthrough
	case 8:
		fmt.Printf("show fallthrought\n")
	default:
		fmt.Printf("default\n")
	}

	// switch без выражения(работает по сути как if)
	age := 19
	switch {
	case age < 18:
		fmt.Println("age < 18")
	case age < 65:
		fmt.Println("age < 65")
	default:
		fmt.Println("age >= 65")
	}

	// type switch
	var value interface{} = "Dianich"
	switch v := value.(type) {
	case string:
		fmt.Println("Строка", v)
	case int:
		fmt.Println("Число", v)
	}
}
