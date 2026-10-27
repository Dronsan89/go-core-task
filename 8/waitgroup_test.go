package main

import (
	"sync/atomic"
	"testing"
	"time"
)

const waitGroupTestTimeout = 2 * time.Second

func TestWaitGroupWaitsForDone(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(1)

	started := make(chan struct{})
	returned := make(chan struct{})

	go func() {
		close(started)
		wg.Wait()
		close(returned)
	}()

	<-started

	select {
	case <-returned:
		t.Fatal("Wait завершился до вызова Done")
	case <-time.After(50 * time.Millisecond):
	}

	wg.Done()
	waitForSignal(t, returned, "Wait не завершился после Done")
}

func TestWaitGroupMultipleGoroutines(t *testing.T) {
	const workers = 100

	wg := NewWaitGroup()
	var completed atomic.Int64

	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			completed.Add(1)
		}()
	}

	waitForWaitGroup(t, wg)

	if got := completed.Load(); got != workers {
		t.Errorf("завершилось %d горутин, ожидалось %d", got, workers)
	}
}

func TestWaitGroupDoneBeforeWait(t *testing.T) {
	wg := NewWaitGroup()

	wg.Add(2)
	wg.Done()
	wg.Done()

	waitForWaitGroup(t, wg)
}

func TestWaitGroupAddZero(t *testing.T) {
	wg := NewWaitGroup()

	wg.Add(0)

	waitForWaitGroup(t, wg)
}

func TestWaitGroupMultipleWaiters(t *testing.T) {
	const waiters = 10

	wg := NewWaitGroup()
	wg.Add(1)

	ready := make(chan struct{}, waiters)
	done := make(chan struct{}, waiters)

	for i := 0; i < waiters; i++ {
		go func() {
			ready <- struct{}{}
			wg.Wait()
			done <- struct{}{}
		}()
	}

	for i := 0; i < waiters; i++ {
		<-ready
	}

	wg.Done()

	for i := 0; i < waiters; i++ {
		waitForSignal(t, done, "не все ожидающие Wait были разблокированы")
	}
}

func TestWaitGroupReuse(t *testing.T) {
	wg := NewWaitGroup()

	wg.Add(1)
	wg.Done()
	waitForWaitGroup(t, wg)

	wg.Add(1)
	wg.Done()
	waitForWaitGroup(t, wg)
}

func TestWaitGroupNegativeCounterPanics(t *testing.T) {
	wg := NewWaitGroup()

	var recovered any
	func() {
		defer func() {
			recovered = recover()
		}()

		wg.Add(-1)
	}()

	if recovered == nil {
		t.Fatal("ожидалась паника при отрицательном счётчике")
	}

	// Проверяем, что семафор освобождён после паники.
	wg.Add(1)
	wg.Done()
	waitForWaitGroup(t, wg)
}

func waitForWaitGroup(t *testing.T, wg *WaitGroup) {
	t.Helper()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	waitForSignal(t, done, "Wait не завершился")
}

func waitForSignal(t *testing.T, ch <-chan struct{}, message string) {
	t.Helper()

	select {
	case <-ch:
	case <-time.After(waitGroupTestTimeout):
		t.Fatal(message)
	}
}
