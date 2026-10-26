package main

import (
	"slices"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name       string
		first      []int
		second     []int
		wantHas    bool
		wantValues []int
	}{
		{
			name:       "пример из задания",
			first:      []int{65, 3, 58, 678, 64},
			second:     []int{64, 2, 3, 43},
			wantHas:    true,
			wantValues: []int{64, 3},
		},
		{
			name:       "нет пересечений",
			first:      []int{1, 2},
			second:     []int{3, 4},
			wantHas:    false,
			wantValues: []int{},
		},
		{
			name:       "один общий элемент",
			first:      []int{1},
			second:     []int{1},
			wantHas:    true,
			wantValues: []int{1},
		},
		{
			name:       "порядок второго слайса",
			first:      []int{3, 64},
			second:     []int{64, 3},
			wantHas:    true,
			wantValues: []int{64, 3},
		},
		{
			name:       "пустой первый слайс",
			first:      []int{},
			second:     []int{1, 2},
			wantHas:    false,
			wantValues: []int{},
		},
		{
			name:       "пустой второй слайс",
			first:      []int{1, 2},
			second:     []int{},
			wantHas:    false,
			wantValues: []int{},
		},
		{
			name:       "оба слайса nil",
			first:      nil,
			second:     nil,
			wantHas:    false,
			wantValues: []int{},
		},
		{
			name:       "повторы в обоих слайсах не дублируют результат",
			first:      []int{1, 1, 2},
			second:     []int{1, 2, 2},
			wantHas:    true,
			wantValues: []int{1, 2},
		},
		{
			name:       "повторы во втором слайсе",
			first:      []int{1, 2},
			second:     []int{1, 1, 2, 2},
			wantHas:    true,
			wantValues: []int{1, 2},
		},
		{
			name:       "повторы в первом слайсе",
			first:      []int{1, 1, 2},
			second:     []int{2, 1},
			wantHas:    true,
			wantValues: []int{2, 1},
		},
		{
			name:       "отрицательные числа и ноль",
			first:      []int{-1, 0, 5},
			second:     []int{0, -1, 7},
			wantHas:    true,
			wantValues: []int{0, -1},
		},
		{
			name:       "все элементы общие",
			first:      []int{1, 2, 3},
			second:     []int{2, 3, 1},
			wantHas:    true,
			wantValues: []int{2, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHas, gotValues := intersection(tt.first, tt.second)

			if gotHas != tt.wantHas {
				t.Errorf("intersection(%v, %v): bool = %t, ожидалось %t",
					tt.first, tt.second, gotHas, tt.wantHas)
			}
			if !slices.Equal(gotValues, tt.wantValues) {
				t.Errorf("intersection(%v, %v) = %v, ожидалось %v",
					tt.first, tt.second, gotValues, tt.wantValues)
			}
		})
	}
}
