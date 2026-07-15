package main

import "fmt"

func main() {

	// int
	var numInt int = 2
	fmt.Printf("int numInt = %d, type of numInt = %T\n", numInt, numInt)

	// int8
	var numInt8 int8 = 2
	fmt.Printf("int8 numInt8 = %d, type of numInt8 = %T\n", numInt8, numInt8)

	// int16
	var numInt16 int16 = 2
	fmt.Printf("int16 numInt16 = %d, type of numInt16 = %T\n", numInt16, numInt16)

	// int32
	var numInt32 int32 = 2
	fmt.Printf("int32 numInt32 = %d, type of numInt32 = %T\n", numInt32, numInt32)

	// int64
	var numInt64 int64 = 2
	fmt.Printf("int64 numInt64 = %d, type of numInt64 = %T\n", numInt64, numInt64)

	// uint
	var numUint uint = 2
	fmt.Printf("uint numUint = %d, type of numUint = %T\n", numUint, numUint)

	// uint8
	var numUint8 uint8 = 2
	fmt.Printf("uint8 numUint8 = %d, type of numUint8 = %T\n", numUint8, numUint8)

	// uint16
	var numUint16 uint16 = 2
	fmt.Printf("uint16 numUint16 = %d, type of numUint16 = %T\n", numUint16, numUint16)

	// uint32
	var numUint32 uint32 = 2
	fmt.Printf("uint32 numUint32 = %d, type of numUint32 = %T\n", numUint32, numUint32)

	// uint64
	var numUint64 uint64 = 2
	fmt.Printf("uint64 numUint64 = %d, type of numUint64 = %T\n", numUint64, numUint64)

	// float32
	var numFloat32 float32 = 2.222
	fmt.Printf("float32 numFloat32 = %f, type of numFloat32 = %T\n", numFloat32, numFloat32)

	// float64
	var numFloat64 float64 = 2.222
	fmt.Printf("float64 numFloat64 = %f, type of numFloat64 = %T\n", numFloat64, numFloat64)

	// complex64
	var numComplex64 complex64 = 1 + 2i
	fmt.Printf("complex64 numComplex64 = %v, type of numComplex64 = %T\n", numComplex64, numComplex64)

	// complex128
	var numComplex128 complex128 = 1 + 2i
	fmt.Printf("complex128 numComplex128 = %v, type of numComplex128 = %T\n", numComplex128, numComplex128)

	// bool
	var varBool bool = true
	fmt.Printf("bool varBool = %t, type of varBool = %T\n", varBool, varBool)

	// string
	var varString string = "string"
	fmt.Printf("string varString = %s, type of varString = %T\n", varString, varString)

	// rune
	var varRune rune = 'r'
	fmt.Printf("rune varRune = %c, type of varRune = %T\n", varRune, varRune)

	// byte
	var varByte byte = 'D'
	fmt.Printf("byte varByte = %c, type of varByte = %T\n", varByte, varByte)

	//// Какие у нас есть варианты объявления переменных

	// Полное объявление
	var fullDeclareVar int
	fmt.Printf("fullDeclareVar =  %d\n", fullDeclareVar)

	// С инициализацией
	var fullDeclareIVar int = 2
	fmt.Printf("fullDeclareIVar = %d\n", fullDeclareIVar)

	// С автоматическим определением типа на основе значения
	var aVar = 2
	fmt.Printf("aVar = %d, aVar type = %T\n", aVar, aVar)

	// Короткое объявление
	sVar := 2
	fmt.Printf("sVar = %d, sVar type = %T\n", sVar, sVar)

	// Групповое объявление вариант1
	var (
		a int = 222
		b int = 2222
	)
	fmt.Printf("a = %d, a type = %T\n", a, a)
	fmt.Printf("b = %d, b type = %T\n", b, b)

	// Групповое объявление вариант2
	var a2, b2 = 222, 2222
	fmt.Printf("a2 = %d, a2 type = %T\n", a2, a2)
	fmt.Printf("b2 = %d, b2 type = %T\n", b2, b2)

	// Групповое объявление вариант3
	a3, b3 := 222, 2222
	fmt.Printf("a3 = %d, a3 type = %T\n", a3, a3)
	fmt.Printf("b3 = %d, b3 type = %T\n", b3, b3)

	// Константы
	const PI = 3.14
	fmt.Printf("PI = %f, PI type = %T\n", PI, PI)

	// Iota
	const (
		One = iota + 1
		Two
		Three
	)

	fmt.Printf("One = %d, One type = %T\n", One, One)
	fmt.Printf("Two = %d, Two type = %T\n", Two, Two)
	fmt.Printf("Three = %d, Three type = %T\n", Three, Three)

	const (
		KB = 1 << (10 * iota)
		MB
		GB
	)
	fmt.Printf("KB = %d, KB type = %T\n", KB, KB)
	fmt.Printf("MB = %d, MB type = %T\n", MB, MB)
	fmt.Printf("GB = %d, GB type = %T\n", GB, GB)

}
