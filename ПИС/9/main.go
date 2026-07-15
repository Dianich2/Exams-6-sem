package main

import "fmt"

func ChangeMap(tMap map[int]int) {
	tMap[22] = 22
}

func main() {
	// объявление map(но тут типо она nil)
	// и еще покажем сразу len
	var tests map[string]string
	fmt.Printf("tests = %v, type = %T, len = %d\n", tests, tests, len(tests))
	fmt.Printf("tests == nil = %t\n", tests == nil)

	// через make(наилучший вариант)
	tests2 := make(map[string]string)
	fmt.Printf("tests2 = %v, type = %T, len = %d\n", tests2, tests2, len(tests2))
	fmt.Printf("tests2 == nil = %t\n", tests2 == nil)

	// создание со значениями
	tests3 := map[string]string{
		"test1": "1 + 1",
		"test2": "2 + 2",
	}
	fmt.Printf("tests3 = %v, type = %T, len = %d\n", tests3, tests3, len(tests3))
	fmt.Printf("tests3 == nil = %t\n", tests3 == nil)

	// добавление значения
	tests3["test3"] = "3 + 3"
	fmt.Printf("tests3 = %v, type = %T, len = %d\n", tests3, tests3, len(tests3))

	// получение значения
	curTest := tests3["test2"]
	fmt.Printf("curTest = %s\n", curTest)

	// попытка получить значение, если ключа нету(в нашем случае вернет пустую строку)
	// а в целом возвращается Zero Value для того типа, который у нас для значения
	emptyTest := tests3["test22"]
	fmt.Printf("emptyTest = %s\n", emptyTest)

	// как понять существует ли ключ?
	val1, ok1 := tests3["test2"]
	fmt.Printf("val1 = %s, ok1 = %t\n", val1, ok1)

	val2, ok2 := tests3["test22"]
	fmt.Printf("val2 = %s, ok2 = %t\n", val2, ok2)

	// изменение элемента
	fmt.Printf("tests3 = %v, type = %T, len = %d\n", tests3, tests3, len(tests3))
	tests3["test2"] = "22 + 22"
	fmt.Printf("tests3 = %v, type = %T, len = %d\n", tests3, tests3, len(tests3))

	// удаление элемента(если ключа нету, то ошибки не будет)
	fmt.Printf("tests3 = %v, type = %T, len = %d\n", tests3, tests3, len(tests3))
	delete(tests3, "test2")
	fmt.Printf("tests3 = %v, type = %T, len = %d\n", tests3, tests3, len(tests3))

	// перебор map с помощью range(важно!!! порядок обхода не гарантируется)
	tests3["test2"] = "22 + 22"
	fmt.Println("for ... range")
	for key, val := range tests3 {
		fmt.Printf("key = %s, val = %s\n", key, val)
	}

	// важно тоже сказать, что ключом могут быть только сравнимые типы
	// вот так нельзя
	//no := map[[]int]string
	//no2 := map[map[string]string]string

	// при попытке копировать map ссылается на одну и ту же память
	map1 := make(map[int]int)
	map1[1] = 1
	map2 := map1
	fmt.Printf("map1 = %v, type = %T, len = %d\n", map1, map1, len(map1))
	fmt.Printf("map2 = %v, type = %T, len = %d\n", map2, map2, len(map2))
	map1[2] = 2
	map2[3] = 3
	fmt.Printf("map1 = %v, type = %T, len = %d\n", map1, map1, len(map1))
	fmt.Printf("map2 = %v, type = %T, len = %d\n", map2, map2, len(map2))

	// передача в функцию(по умолчанию по ссылке; функция выше)
	fmt.Printf("map1 = %v, type = %T, len = %d\n", map1, map1, len(map1))
	ChangeMap(map1)
	fmt.Printf("map1 = %v, type = %T, len = %d\n", map1, map1, len(map1))

	// вложенные map
	lMap := make(map[int]map[string]string)
	lMap[1] = make(map[string]string)
	lMap[2] = make(map[string]string)
	lMap[1]["Dianich"] = "test1"
	lMap[1]["Dianich2"] = "test2"
	lMap[2]["Petya"] = "test222"
	fmt.Printf("lMap = %v, type = %T, len = %d\n", lMap, lMap, len(lMap))
}
