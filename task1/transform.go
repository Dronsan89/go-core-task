package main

import (
	"crypto/sha256"
	"fmt"
)

// Salt — добавляемая в середину среза рун.
const Salt = "go-2024"

// PrintTypes возвращает строки с типом и значением каждой переменной.
func PrintTypes(vars ...interface{}) []string {
	result := make([]string, 0, len(vars))
	for _, v := range vars {
		result = append(result, fmt.Sprintf("Тип: %T, Значение: %v", v, v))
	}
	return result
}

// ConvertToStrings преобразует переменные в строковый тип.
func ConvertToStrings(vars ...interface{}) []string {
	result := make([]string, 0, len(vars))
	for _, v := range vars {
		result = append(result, fmt.Sprintf("%v", v))
	}
	return result
}

// ConcatStrings объединяет срез строк в одну строку.
func ConcatStrings(parts []string) string {
	result := ""
	for _, p := range parts {
		result += p
	}
	return result
}

// StringToRunes преобразует строку в срез рун.
func StringToRunes(s string) []rune {
	return []rune(s)
}

// InsertSalt вставляет соль в середину среза рун.
func InsertSalt(runes []rune, salt string) []rune {
	saltRunes := []rune(salt)
	mid := len(runes) / 2
	result := make([]rune, 0, len(runes)+len(saltRunes))
	result = append(result, runes[:mid]...)
	result = append(result, saltRunes...)
	result = append(result, runes[mid:]...)
	return result
}

// HashRunes вычисляет SHA256 от среза рун (в виде строки).
func HashRunes(runes []rune) string {
	hash := sha256.Sum256([]byte(string(runes)))
	return fmt.Sprintf("%x", hash)
}

// ProcessVars выполняет весь пайплайн: приводит переменные к строкам,
// объединяет, добавляет соль в середину и возвращает хэш SHA256.
func ProcessVars(vars ...interface{}) string {
	strs := ConvertToStrings(vars...)
	joined := ConcatStrings(strs)
	runes := StringToRunes(joined)
	withSalt := InsertSalt(runes, Salt)
	return HashRunes(withSalt)
}
