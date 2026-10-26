package stringintmap

import (
	"reflect"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	m := New()

	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("", 0)

	tests := []struct {
		key       string
		wantValue int
		wantOK    bool
	}{
		{key: "one", wantValue: 1, wantOK: true},
		{key: "two", wantValue: 2, wantOK: true},
		{key: "", wantValue: 0, wantOK: true},
		{key: "missing", wantValue: 0, wantOK: false},
	}

	for _, tt := range tests {
		value, ok := m.Get(tt.key)
		if value != tt.wantValue || ok != tt.wantOK {
			t.Errorf("Get(%q) = (%d, %t), ожидалось (%d, %t)",
				tt.key, value, ok, tt.wantValue, tt.wantOK)
		}
	}
}

func TestAddOverwritesExistingKey(t *testing.T) {
	m := New()
	m.Add("key", 1)
	m.Add("key", 2)

	value, ok := m.Get("key")
	if !ok {
		t.Fatal("ключ должен существовать после повторного добавления")
	}
	if value != 2 {
		t.Errorf("Get(\"key\") = %d, ожидалось 2", value)
	}
}

func TestExists(t *testing.T) {
	m := New()
	m.Add("present", 10)

	if !m.Exists("present") {
		t.Error("Exists(\"present\") = false, ожидалось true")
	}
	if m.Exists("missing") {
		t.Error("Exists(\"missing\") = true, ожидалось false")
	}
}

func TestRemove(t *testing.T) {
	m := New()
	m.Add("a", 1)
	m.Add("b", 2)

	m.Remove("a")

	if m.Exists("a") {
		t.Error("ключ \"a\" должен быть удалён")
	}
	if value, ok := m.Get("a"); ok || value != 0 {
		t.Errorf("Get(\"a\") = (%d, %t), ожидалось (0, false)", value, ok)
	}
	if !m.Exists("b") {
		t.Error("удаление \"a\" не должно затрагивать ключ \"b\"")
	}

	// Удаление отсутствующего ключа не должно вызывать панику.
	m.Remove("missing")
}

func TestCopy(t *testing.T) {
	m := New()
	m.Add("a", 1)
	m.Add("b", 2)

	got := m.Copy()
	want := map[string]int{"a": 1, "b": 2}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Copy() = %v, ожидалось %v", got, want)
	}
}

func TestCopyIsIndependent(t *testing.T) {
	m := New()
	m.Add("a", 1)
	m.Add("b", 2)

	copied := m.Copy()
	copied["a"] = 100
	delete(copied, "b")
	copied["c"] = 3

	if value, ok := m.Get("a"); !ok || value != 1 {
		t.Errorf("изменение копии затронуло оригинал: Get(\"a\") = (%d, %t)", value, ok)
	}
	if !m.Exists("b") {
		t.Error("удаление из копии затронуло оригинал")
	}
	if m.Exists("c") {
		t.Error("добавление в копию затронуло оригинал")
	}
}
