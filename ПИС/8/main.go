package main

import "fmt"

func TryChangeSliceByValue(s []int) {
	s[0] = 9
	s[1] = 9
	s = append(s, 5)
}

func TryChangeSliceByPtr(s *[]int) {
	(*s)[0] = 2222
	(*s)[1] = 2222
	*s = append(*s, 5)
}

func main() {
	// объявление способ1
	nums := []int{2, 22, 222}
	fmt.Printf("nums = %v, type = %T\n", nums, nums)

	// объявление через var(будет nil)
	var nums2 []int
	fmt.Printf("nums2 = %v, type = %T\n", nums2, nums2)
	fmt.Printf("nums2 == nil = %t\n", nums2 == nil)

	// объявление через make (самый лучший способ)
	// можно сразу задать длину
	nums3 := make([]int, 2)
	fmt.Printf("nums3 = %v, type = %T, len = %d, cap = %d\n", nums3, nums3, len(nums3), cap(nums3))

	// а можно еще и capacity
	nums4 := make([]int, 2, 22)
	fmt.Printf("nums4 = %v, type = %T, len = %d, cap = %d\n", nums4, nums4, len(nums4), cap(nums4))

	// добавление элементов (append)
	nums4 = append(nums4, 2)
	fmt.Printf("nums4 = %v, type = %T, len = %d, cap = %d\n", nums4, nums4, len(nums4), cap(nums4))

	// и сразу несколько добавим
	nums4 = append(nums4, 222, 222, 222)
	fmt.Printf("nums4 = %v, type = %T, len = %d, cap = %d\n", nums4, nums4, len(nums4), cap(nums4))

	// создадим срез на основе массива(тут cap = Длина массива - Индекс начала среза)
	arr := [5]int{2, 22, 222, 2222, 22222}
	nums5 := arr[2:4]
	fmt.Printf("arr = %v, type = %T, len = %d\n", arr, arr, len(arr))
	fmt.Printf("nums5 = %v, type = %T, len = %d, cap = %d\n", nums5, nums5, len(nums5), cap(nums5))

	// покажем, что массив и срез используют одну память
	nums5[0] = 0
	nums5 = append(nums5, 4)
	arr[3] = 9
	fmt.Printf("arr = %v, type = %T, len = %d\n", arr, arr, len(arr))
	fmt.Printf("nums5 = %v, type = %T, len = %d, cap = %d\n", nums5, nums5, len(nums5), cap(nums5))

	// копирование срезов
	// тут по сути они будут ссылаться на один и тот же участок памяти
	// покажем это тоже
	a := []int{22, 22, 22}
	b := a
	fmt.Printf("a = %v, address a[0] = %p, type = %T, len = %d, cap = %d\n", a, &a[0], a, len(a), cap(a))
	fmt.Printf("b = %v, address b[0] = %p, type = %T, len = %d, cap = %d\n", b, &b[0], b, len(b), cap(b))
	a[0] = 9
	b[1] = 5
	fmt.Printf("a = %v, address a[0] = %p, type = %T, len = %d, cap = %d\n", a, &a[0], a, len(a), cap(a))
	fmt.Printf("b = %v, address b[0] = %p, type = %T, len = %d, cap = %d\n", b, &b[0], b, len(b), cap(b))

	// а вот способ, если мы хотим сделать срезы независимыми
	c := []int{2, 2, 2}
	d := make([]int, len(a))
	copy(d, c)
	fmt.Printf("c = %v, address c[0] = %p, type = %T, len = %d, cap = %d\n", c, &c[0], c, len(c), cap(c))
	fmt.Printf("d = %v, address d[0] = %p, type = %T, len = %d, cap = %d\n", d, &d[0], d, len(d), cap(d))
	c[0] = 9
	d[1] = 5
	fmt.Printf("c = %v, address c[0] = %p, type = %T, len = %d, cap = %d\n", c, &c[0], c, len(c), cap(c))
	fmt.Printf("d = %v, address d[0] = %p, type = %T, len = %d, cap = %d\n", d, &d[0], d, len(d), cap(d))

	// удаление элемента(встроенной функции нет)
	nums6 := []int{2, 22, 222}
	fmt.Printf("nums6 = %v, type = %T, len = %d, cap = %d\n", nums6, nums6, len(nums6), cap(nums6))
	nums6 = append(nums6[:1], nums6[2:]...)
	fmt.Printf("nums6 = %v, type = %T, len = %d, cap = %d\n", nums6, nums6, len(nums6), cap(nums6))

	// перебор среза через цикл for
	fmt.Println("for")
	for i := 0; i < len(nums6); i++ {
		fmt.Printf("nums6[%d] = %d\n", i, nums6[i])
	}

	// перебор с помощью цикла for с использованием range
	fmt.Println("for ... range")
	for idx, val := range nums6 {
		fmt.Printf("nums6[%d] = %d\n", idx, val)
	}

	// передача срезов в функцию(по умолчанию по ссылке(но не все); функция выше)
	// тут важно рассказать про заголовок, разница будет в данном случае
	s := []int{22, 22}
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
	TryChangeSliceByValue(s)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
	TryChangeSliceByPtr(&s)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
