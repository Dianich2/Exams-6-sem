package main

import (
	"fmt"
	"strings"
)

func main() {
	// объявление полностью
	var str1 string = "Dianich"
	fmt.Printf("str1 = %s, type = %T\n", str1, str1)

	// с автоматическим выводом типа
	var str2 = "Dianich"
	fmt.Printf("str2 = %s, type = %T\n", str2, str2)

	// краткая запись
	str3 := "Dianich"
	fmt.Printf("str3 = %s, type = %T\n", str3, str3)

	//// строковые литералы
	// есть обычная строка, тут нужно вводить специальные символы(\n, \r, \t и т.д.)
	baseString := "Dianich"
	fmt.Printf("baseString = %s, type = %T\n", baseString, baseString)

	// есть сырые строки(raw-строки), тут вот как есть строка, так и будет
	rawString := `Hello
	Golang!!!`
	fmt.Printf("rawString = %s, type = %T\n", rawString, rawString)

	// Длина строки(важно, что она показывается в байтах)
	str4 := "Dianich"
	fmt.Printf("str4 = %s, type = %T, len = %d\n", str4, str4, len(str4))

	str5 := "Дианыч"
	fmt.Printf("str5 = %s, type = %T, len = %d\n", str5, str5, len(str5))

	// доступ к символам строки
	fmt.Printf("str4[0] = %c\n", str4[0])

	//// но чаще всего для работы со строками используют []rune

	// перебор строки с помощью for
	fmt.Println("for")
	for i := 0; i < len(str4); i++ {
		fmt.Printf("str4[%d] = %c\n", i, str4[i])
	}

	// перебор строки с помощью for ... range
	fmt.Println("for ... range")
	for idx, sym := range str4 {
		fmt.Printf("str4[%d] = %c\n", idx, sym)
	}

	// строки не изменяются!!!

	// конкатенация строк
	str6 := str4 + str4
	fmt.Printf("str6 = %s, type = %T, len = %d\n", str6, str6, len(str6))

	// сравнение строк
	fmt.Printf("str4 > str6 = %t\nstr4 < str6 = %t\nstr4 >= str6 = %t\nstr4 <= str6 = %t\nstr4 == str6 = %t",
		str4 > str6, str4 < str6, str4 >= str6, str4 <= str6, str4 == str6)

	// пакет strings
	// Contains
	fmt.Printf("str6 contains str4 = %t\n", strings.Contains(str6, str4))

	// HasPrefix
	fmt.Printf("str6 hasPrefix str4 = %t\n", strings.HasPrefix(str6, str4))

	// HasSuffix
	fmt.Printf("str6 hasSuffix str4 = %t\n", strings.HasSuffix(str6, str4))

	// ToUpper
	fmt.Printf("str6 toUpper = %s\n", strings.ToUpper(str6))

	// ToLower
	fmt.Printf("str6 toLower = %s\n", strings.ToLower(str6))

	// ReplaceAll
	fmt.Printf("str6 replace i - u = %s\n", strings.ReplaceAll(str6, "i", "u"))

	// Split
	fmt.Printf("a, b, c split = %v\n", strings.Split("a, b, c", ", "))

	// Join
	fmt.Printf("a, b, c join = %s\n", strings.Join([]string{"a", "b", "c"}, ", "))

	// TrimSpace
	fmt.Printf("trimSpace = %s\n", strings.TrimSpace("    Dianich    "))

	// Index
	fmt.Printf("index = %d\n", strings.Index(str6, "i"))

	// Рассказать про strings.Builder и []byte
}
