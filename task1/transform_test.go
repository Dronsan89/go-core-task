package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestPrintTypes(t *testing.T) {
	result := PrintTypes(42, 3.14, "Golang", true, complex64(1+2i))
	if len(result) != 5 {
		t.Fatalf("ожидалось 5 строк, получено %d", len(result))
	}
	// Проверяем, что в выводе фигурируют ожидаемые типы
	expectedTypes := []string{"int", "float64", "string", "bool", "complex64"}
	for i, exp := range expectedTypes {
		if !contains(result[i], exp) {
			t.Errorf("строка %d не содержит тип %q: %q", i, exp, result[i])
		}
	}
}

func TestConvertToStrings(t *testing.T) {
	cases := []struct {
		in       interface{}
		expected string
	}{
		{42, "42"},
		{3.14, "3.14"},
		{"Golang", "Golang"},
		{true, "true"},
		{complex64(1 + 2i), "(1+2i)"},
	}
	for _, c := range cases {
		got := ConvertToStrings(c.in)
		if len(got) > 1 || got[0] != c.expected {
			t.Errorf("ConvertToStrings(%v) = %v, ожидалось [%q]", c.in, got, c.expected)
		}
	}
}

func TestConcatStrings(t *testing.T) {
	got := ConcatStrings([]string{"a", "b", "c"})
	if got != "abc" {
		t.Errorf("ConcatStrings = %q, ожидалось %q", got, "abc")
	}
	if ConcatStrings(nil) != "" {
		t.Errorf("ConcatStrings(nil) должен быть пустой строкой")
	}
}

func TestStringToRunes(t *testing.T) {
	got := StringToRunes("Привет")
	if len(got) != 6 {
		t.Errorf("ожидалось 6 рун, получено %d", len(got))
	}
	if got[0] != 'П' {
		t.Errorf("первая руна = %q, ожидалось 'П'", got[0])
	}
}

func TestInsertSalt(t *testing.T) {
	runes := []rune("abcdef")
	got := InsertSalt(runes, "go-2024")
	expected := "abcgo-2024def"
	if string(got) != expected {
		t.Errorf("InsertSalt = %q, ожидалось %q", string(got), expected)
	}

	// Проверка на пустом срезе: соль должна быть вставлена
	empty := InsertSalt([]rune{}, "salt")
	if string(empty) != "salt" {
		t.Errorf("InsertSalt(пусто) = %q, ожидалось %q", string(empty), "salt")
	}

	// Проверка на нечётной длине
	odd := InsertSalt([]rune("abc"), "X")
	if string(odd) != "aXbc" {
		t.Errorf("InsertSalt(abc) = %q, ожидалось %q", string(odd), "aXbc")
	}
}

func TestHashRunes(t *testing.T) {
	input := []rune("test")
	expected := fmt.Sprintf("%x", sha256.Sum256([]byte("test")))
	got := HashRunes(input)
	if got != expected {
		t.Errorf("HashRunes = %q, ожидалось %q", got, expected)
	}
	if len(got) != 64 {
		t.Errorf("SHA256 hex должен иметь длину 64, получено %d", len(got))
	}
}

func TestProcessVars(t *testing.T) {
	vars := []interface{}{42, 3.14, "Golang", true, complex64(1 + 2i)}
	got := ProcessVars(vars...)

	// Воспроизводим ожидаемый результат вручную
	joined := "423.14Golangtrue(1+2i)"
	runes := []rune(joined)
	withSalt := InsertSalt(runes, Salt)
	expected := fmt.Sprintf("%x", sha256.Sum256([]byte(string(withSalt))))

	if got != expected {
		t.Errorf("ProcessVars = %q, ожидалось %q", got, expected)
	}
	if len(got) != 64 {
		t.Errorf("хэш должен иметь длину 64, получено %d", len(got))
	}
}

// contains — вспомогательная функция проверки подстроки.
func contains(s, substr string) bool {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
