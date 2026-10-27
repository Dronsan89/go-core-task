package main

// semaphore — бинарный семафор на основе буферизированного канала.
type semaphore chan struct{}

func newSemaphore() semaphore {
	return make(semaphore, 1)
}

func (s semaphore) acquire() {
	s <- struct{}{}
}

func (s semaphore) release() {
	<-s
}

// WaitGroup — кастомная замена sync.WaitGroup, построенная на семафоре.
// Создавайте её через NewWaitGroup.
type WaitGroup struct {
	sem   semaphore
	done  chan struct{}
	count int
}

// NewWaitGroup создаёт новый WaitGroup.
func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		sem:  newSemaphore(),
		done: make(chan struct{}),
	}
}

// Add изменяет счётчик WaitGroup на delta.
// При отрицательном значении счётчика возникает паника.
func (wg *WaitGroup) Add(delta int) {
	wg.sem.acquire()
	defer wg.sem.release()

	oldCount := wg.count
	newCount := oldCount + delta
	if newCount < 0 {
		panic("WaitGroup: отрицательный счётчик")
	}

	// Начинается новое поколение работы после предыдущего Wait.
	if oldCount == 0 && delta > 0 {
		wg.done = make(chan struct{})
	}

	wg.count = newCount

	if oldCount > 0 && newCount == 0 {
		close(wg.done)
	}
}

// Done уменьшает счётчик на единицу.
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

// Wait блокируется, пока счётчик не станет равен нулю.
func (wg *WaitGroup) Wait() {
	wg.sem.acquire()

	if wg.count == 0 {
		wg.sem.release()
		return
	}

	done := wg.done
	wg.sem.release()

	<-done
}
