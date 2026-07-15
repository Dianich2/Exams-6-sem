package main

import "fmt"

func ChangeArrayByValue(arr [2]int) {
	arr[0] = 2222
	arr[1] = 2222
}

func ChangeArrayByPtr(arr *[2]int) {
	arr[0] = 2222
	arr[1] = 2222
}

func main() {
	// объявление (тут покажем как в целом объявить ну и Zero Value)
	var nums [2]int
	fmt.Printf("nums = %v\n", nums)

	// полная инициализация
	var nums2 [2]int = [2]int{2, 22}
	fmt.Printf("nums2 = %v\n", nums2)

	// сокращенная запись
	nums3 := [2]int{2, 22}
	fmt.Printf("nums3 = %v\n", nums3)

	// автоматическое определение размера
	// плюс сразу покажем как посмотреть длину
	nums4 := [...]int{2, 22}
	fmt.Printf("nums4 = %v, len = %d\n", nums4, len(nums4))

	// перебор с помощью цикла for
	fmt.Println("for")
	for i := 0; i < len(nums); i++ {
		fmt.Printf("nums[%d] = %d\n", i, nums[i])
	}

	// перебор с помощью цикла for с использованием range
	// (тут в принципе если нам не нужен индекс, то можно скипать с помощью _ вместо idx)
	fmt.Println("for ... range")
	for idx, val := range nums {
		fmt.Printf("nums[%d] = %d\n", idx, val)
	}

	// важно сказать, что размер это часть типа
	// ну и тут просто примерчик, что нельзя приравнивать такие массивы
	var a [2]int
	var b [22]int
	fmt.Printf("a = %v, len = %d, type = %T\n", a, len(a), a)
	fmt.Printf("b = %v, len = %d, type = %T\n", b, len(b), b)
	// a = b      получаем ошибку компиляции

	// сравнение массивов(опять же должен быть один и тот же тип)
	// массивы равны, если равны типы и все значения
	x := [2]int{2, 22}
	y := [2]int{2, 222}
	fmt.Printf("x == y = %t\n", x == y)

	// многомерные массивы
	matrix := [2][2]int{
		{2, 2},
		{22, 22},
	}
	fmt.Printf("matrix = %v, len = %d, type = %T\n", matrix, len(matrix), matrix)

	// передача массивов в функцию(по умолчанию по значению; функция выше)
	var arr1 [2]int
	fmt.Printf("arr1 = %v\n", arr1)
	ChangeArrayByValue(arr1)
	fmt.Printf("arr1 = %v\n", arr1)
	ChangeArrayByPtr(&arr1)
	fmt.Printf("arr1 = %v\n", arr1)
}
