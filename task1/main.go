package main

import "fmt"

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	vars := []interface{}{
		numDecimal, numOctal, numHexadecimal,
		pi, name, isActive, complexNum,
	}

	fmt.Println("=== Типы переменных ===")
	for _, line := range PrintTypes(vars...) {
		fmt.Println(line)
	}

	fmt.Println("\n=== Строковые представления ===")
	strs := ConvertToStrings(vars...)
	for i, s := range strs {
		fmt.Printf("%d: %q\n", i, s)
	}

	joined := ConcatStrings(strs)
	fmt.Printf("\n=== Объединённая строка ===\n%s\n", joined)

	runes := StringToRunes(joined)
	fmt.Printf("\n=== Срез рун (длина: %d) ===\n%v\n", len(runes), runes)

	withSalt := InsertSalt(runes, Salt)
	fmt.Printf("\n=== С солью \"%s\" в середине ===\n%q\n", Salt, string(withSalt))

	hash := HashRunes(withSalt)
	fmt.Printf("\n=== SHA256 хэш ===\n%s\n", hash)

	fmt.Printf("\n=== Итоговый хэш через ProcessVars ===\n%s\n", ProcessVars(vars...))
}
