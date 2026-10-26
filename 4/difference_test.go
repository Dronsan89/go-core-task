package main

import (
	"slices"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name   string
		first  []string
		second []string
		want   []string
	}{
		{
			name:   "пример из задания",
			first:  []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			second: []string{"banana", "date", "fig"},
			want:   []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:   "второй слайс пустой",
			first:  []string{"a", "b"},
			second: []string{},
			want:   []string{"a", "b"},
		},
		{
			name:   "все элементы первого есть во втором",
			first:  []string{"a", "b"},
			second: []string{"b", "a"},
			want:   []string{},
		},
		{
			name:   "первый слайс пустой",
			first:  []string{},
			second: []string{"a"},
			want:   []string{},
		},
		{
			name:   "оба слайса nil",
			first:  nil,
			second: nil,
			want:   []string{},
		},
		{
			name:   "повторы первого слайса сохраняются",
			first:  []string{"a", "a", "b", "a"},
			second: []string{"b"},
			want:   []string{"a", "a", "a"},
		},
		{
			name:   "повторы во втором слайсе не влияют на результат",
			first:  []string{"a", "b"},
			second: []string{"b", "b", "b"},
			want:   []string{"a"},
		},
		{
			name:   "порядок первого слайса сохраняется",
			first:  []string{"c", "a", "b"},
			second: []string{"a"},
			want:   []string{"c", "b"},
		},
		{
			name:   "пустая строка является обычным элементом",
			first:  []string{"", "a", ""},
			second: []string{"a"},
			want:   []string{"", ""},
		},
		{
			name:   "регистр символов учитывается",
			first:  []string{"Apple", "apple"},
			second: []string{"apple"},
			want:   []string{"Apple"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := difference(tt.first, tt.second)
			if !slices.Equal(got, tt.want) {
				t.Errorf("difference(%v, %v) = %v, ожидалось %v",
					tt.first, tt.second, got, tt.want)
			}
		})
	}
}

func TestDifferenceDoesNotModifyInputs(t *testing.T) {
	first := []string{"apple", "banana", "cherry"}
	second := []string{"banana"}

	difference(first, second)

	if !slices.Equal(first, []string{"apple", "banana", "cherry"}) {
		t.Errorf("первый слайс изменился: %v", first)
	}
	if !slices.Equal(second, []string{"banana"}) {
		t.Errorf("второй слайс изменился: %v", second)
	}
}

func TestDifferenceReturnsIndependentSlice(t *testing.T) {
	first := []string{"a", "b"}
	got := difference(first, nil)

	got[0] = "changed"

	if first[0] != "a" {
		t.Errorf("изменение результата затронуло исходный слайс: %v", first)
	}
}
