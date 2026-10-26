package main

import (
	"reflect"
	"testing"
)

func TestGenerateRandomSlice(t *testing.T) {
	got := generateRandomSlice(10)
	if len(got) != 10 {
		t.Fatalf("ожидалось 10 элементов, получено %d", len(got))
	}

	for i, value := range got {
		if value < 0 || value >= randomLimit {
			t.Errorf("элемент %d = %d, ожидался диапазон [0, %d)", i, value, randomLimit)
		}
	}

	if got := generateRandomSlice(0); len(got) != 0 {
		t.Errorf("generateRandomSlice(0) = %v, ожидался пустой слайс", got)
	}
}

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "смешанные числа", input: []int{1, 2, 3, 4, 5, 6}, want: []int{2, 4, 6}},
		{name: "только нечётные", input: []int{1, 3, 5}, want: []int{}},
		{name: "только чётные", input: []int{0, -2, 8}, want: []int{0, -2, 8}},
		{name: "nil", input: nil, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceExample(%v) = %v, ожидалось %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	want := []int{1, 2, 3, 4}

	got := addElements(input, 4)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addElements(%v, 4) = %v, ожидалось %v", input, got, want)
	}
	if !reflect.DeepEqual(input, []int{1, 2, 3}) {
		t.Errorf("исходный слайс изменился: %v", input)
	}
}

func TestAddElementsReturnsIndependentSlice(t *testing.T) {
	input := make([]int, 2, 10)
	input[0], input[1] = 1, 2

	got := addElements(input, 3)
	got[0] = 100

	if input[0] != 1 {
		t.Errorf("изменение результата затронуло исходный слайс: %v", input)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	got := copySlice(input)

	if !reflect.DeepEqual(got, input) {
		t.Fatalf("copySlice(%v) = %v, ожидалось %v", input, got, input)
	}

	got[0] = 100
	if input[0] != 1 {
		t.Errorf("изменение копии затронуло оригинал: %v", input)
	}

	input[1] = 200
	if got[1] == 200 {
		t.Errorf("изменение оригинала затронуло копию: %v", got)
	}
}

func TestCopySliceNilAndEmpty(t *testing.T) {
	if got := copySlice(nil); got != nil {
		t.Errorf("copySlice(nil) = %v, ожидался nil", got)
	}

	empty := []int{}
	got := copySlice(empty)
	if got == nil {
		t.Error("copySlice([]int{}) вернула nil вместо пустого слайса")
	}
	if len(got) != 0 {
		t.Errorf("copySlice([]int{}) = %v, ожидался пустой слайс", got)
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		index int
		want  []int
	}{
		{name: "середина", input: []int{1, 2, 3, 4}, index: 1, want: []int{1, 3, 4}},
		{name: "первый элемент", input: []int{1, 2, 3}, index: 0, want: []int{2, 3}},
		{name: "последний элемент", input: []int{1, 2, 3}, index: 2, want: []int{1, 2}},
		{name: "единственный элемент", input: []int{1}, index: 0, want: []int{}},
		{name: "пустой слайс", input: []int{}, index: 0, want: []int{}},
		{name: "отрицательный индекс", input: []int{1, 2, 3}, index: -1, want: []int{1, 2, 3}},
		{name: "индекс за верхней границей", input: []int{1, 2, 3}, index: 3, want: []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBefore := make([]int, len(tt.input))
			copy(inputBefore, tt.input)

			got := removeElement(tt.input, tt.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElement(%v, %d) = %v, ожидалось %v", tt.input, tt.index, got, tt.want)
			}
			if !reflect.DeepEqual(tt.input, inputBefore) {
				t.Errorf("исходный слайс изменился: было %v, стало %v", inputBefore, tt.input)
			}
		})
	}
}

func TestRemoveElementReturnsIndependentSlice(t *testing.T) {
	input := []int{1, 2, 3}
	got := removeElement(input, 1)

	got[0] = 100
	if input[0] != 1 {
		t.Errorf("изменение результата затронуло исходный слайс: %v", input)
	}
}
