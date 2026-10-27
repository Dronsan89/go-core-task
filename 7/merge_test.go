package main

import (
	"slices"
	"testing"
	"time"
)

const mergeTestTimeout = 2 * time.Second

func TestMerge(t *testing.T) {
	tests := []struct {
		name   string
		inputs [][]int
		want   []int
	}{
		{
			name:   "два канала",
			inputs: [][]int{{1, 2}, {3, 4}},
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "один канал",
			inputs: [][]int{{5, 6}},
			want:   []int{5, 6},
		},
		{
			name:   "ноль каналов",
			inputs: [][]int{},
			want:   []int{},
		},
		{
			name:   "пустые каналы",
			inputs: [][]int{{}, {}},
			want:   []int{},
		},
		{
			name:   "повторы сохраняются",
			inputs: [][]int{{1, 2}, {2, 3}, {4}},
			want:   []int{1, 2, 2, 3, 4},
		},
		{
			name:   "отрицательные числа и ноль",
			inputs: [][]int{{-1, 0}, {-2}},
			want:   []int{-2, -1, 0},
		},
		{
			name:   "много каналов",
			inputs: [][]int{{10}, {20}, {30}, {40}, {50}},
			want:   []int{10, 20, 30, 40, 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channels := make([]<-chan int, len(tt.inputs))
			for i, values := range tt.inputs {
				channels[i] = channelWithValues(values...)
			}

			got := collectInts(t, merge(channels...))
			slices.Sort(got)

			if !slices.Equal(got, tt.want) {
				t.Errorf("merge() = %v, ожидалось %v", got, tt.want)
			}
		})
	}
}

func TestMergeConsumesChannelsConcurrently(t *testing.T) {
	first := make(chan int)
	second := make(chan int)

	out := merge(first, second)

	go func() {
		second <- 2
	}()

	select {
	case value, ok := <-out:
		if !ok {
			t.Fatal("выходной канал закрылся до получения значения")
		}
		if value != 2 {
			t.Fatalf("получено значение %d, ожидалось 2", value)
		}
	case <-time.After(mergeTestTimeout):
		t.Fatal("merge не читает второй канал, пока первый открыт")
	}

	close(first)
	close(second)

	got := collectInts(t, out)
	if len(got) != 0 {
		t.Errorf("после закрытия каналов получены лишние значения: %v", got)
	}
}

func channelWithValues(values ...int) <-chan int {
	ch := make(chan int, len(values))
	for _, value := range values {
		ch <- value
	}
	close(ch)
	return ch
}

func collectInts(t *testing.T, ch <-chan int) []int {
	t.Helper()

	result := make([]int, 0)
	timeout := time.After(mergeTestTimeout)

	for {
		select {
		case value, ok := <-ch:
			if !ok {
				return result
			}
			result = append(result, value)
		case <-timeout:
			t.Fatal("таймаут: выходной канал не закрылся")
		}
	}
}
