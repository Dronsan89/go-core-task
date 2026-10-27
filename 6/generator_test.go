package main

import (
	"context"
	"testing"
	"time"
)

const testTimeout = 2 * time.Second

func TestRandomInts(t *testing.T) {
	const count = 20

	ch := randomInts(context.Background(), count)
	if cap(ch) != 0 {
		t.Fatalf("cap(ch) = %d, ожидался небуферизированный канал", cap(ch))
	}

	for i := 0; i < count; i++ {
		value, ok := receiveInt(t, ch)
		if !ok {
			t.Fatalf("канал закрылся после %d из %d значений", i, count)
		}
		if value < 0 || value >= randomLimit {
			t.Errorf("значение %d вне диапазона [0, %d)", value, randomLimit)
		}
	}

	select {
	case value, ok := <-ch:
		if ok {
			t.Errorf("получено лишнее значение %d", value)
		}
	case <-time.After(testTimeout):
		t.Fatal("канал не закрылся после отправки всех значений")
	}
}

func TestRandomIntsZeroCount(t *testing.T) {
	ch := randomInts(context.Background(), 0)

	select {
	case value, ok := <-ch:
		if ok {
			t.Errorf("получено значение %d при count = 0", value)
		}
	case <-time.After(testTimeout):
		t.Fatal("канал не закрылся при count = 0")
	}
}

func TestRandomIntsCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const count = 1_000_000

	ch := randomInts(ctx, count)

	for i := 0; i < 5; i++ {
		if _, ok := receiveInt(t, ch); !ok {
			t.Fatalf("канал закрылся до отмены контекста")
		}
	}

	cancel()

	// После отмены производитель может успеть отправить не более одного
	// значения, но канал обязан закрыться без получения всех count значений.
	drained := 0
	deadline := time.After(testTimeout)

	for {
		select {
		case _, ok := <-ch:
			if !ok {
				if drained > 10 {
					t.Fatalf("после отмены получено %d значений, генератор не остановился", drained)
				}
				return
			}
			drained++
		case <-deadline:
			t.Fatal("канал не закрылся после отмены контекста")
		}
	}
}

func receiveInt(t *testing.T, ch <-chan int) (int, bool) {
	t.Helper()

	select {
	case value, ok := <-ch:
		return value, ok
	case <-time.After(testTimeout):
		t.Fatal("таймаут ожидания значения из канала")
		return 0, false
	}
}
