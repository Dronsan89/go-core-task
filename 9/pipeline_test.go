package main

import (
	"slices"
	"testing"
	"time"
)

const pipelineTestTimeout = 2 * time.Second

func TestCube(t *testing.T) {
	tests := []struct {
		value uint8
		want  float64
	}{
		{value: 0, want: 0},
		{value: 1, want: 1},
		{value: 2, want: 8},
		{value: 3, want: 27},
		{value: 10, want: 1000},
		{value: 255, want: 16_581_375},
	}

	for _, tt := range tests {
		if got := cube(tt.value); got != tt.want {
			t.Errorf("cube(%d) = %v, ожидалось %v", tt.value, got, tt.want)
		}
	}
}

func TestCubePipeline(t *testing.T) {
	tests := []struct {
		name  string
		input []uint8
		want  []float64
	}{
		{
			name:  "обычные значения",
			input: []uint8{1, 2, 3, 4},
			want:  []float64{1, 8, 27, 64},
		},
		{
			name:  "порядок сохраняется",
			input: []uint8{3, 1, 2},
			want:  []float64{27, 1, 8},
		},
		{
			name:  "ноль",
			input: []uint8{0},
			want:  []float64{0},
		},
		{
			name:  "верхняя граница uint8",
			input: []uint8{255},
			want:  []float64{16_581_375},
		},
		{
			name:  "пустой вход",
			input: []uint8{},
			want:  []float64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runCubePipeline(t, tt.input)
			if !slices.Equal(got, tt.want) {
				t.Errorf("cubePipeline(%v) = %v, ожидалось %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCubePipelineAllUint8Values(t *testing.T) {
	input := make([]uint8, 0, 256)
	want := make([]float64, 0, 256)

	for value := 0; value <= 255; value++ {
		input = append(input, uint8(value))
		want = append(want, cube(uint8(value)))
	}

	got := runCubePipeline(t, input)
	if !slices.Equal(got, want) {
		t.Error("конвейер вернул неверные значения для полного диапазона uint8")
	}
}

func runCubePipeline(t *testing.T, values []uint8) []float64 {
	t.Helper()

	input := make(chan uint8)
	output := make(chan float64)

	go cubePipeline(input, output)

	go func() {
		defer close(input)

		for _, value := range values {
			input <- value
		}
	}()

	result := make([]float64, 0, len(values))
	timeout := time.After(pipelineTestTimeout)

	for {
		select {
		case value, ok := <-output:
			if !ok {
				return result
			}
			result = append(result, value)
		case <-timeout:
			t.Fatal("таймаут: выходной канал не закрылся")
		}
	}
}
