package concurrency

type Semaphore struct {
	sem chan struct{}
}

func NewSemaphore(capacity int) *Semaphore {
	if capacity <= 0 {
		panic("semaphore size must be positive")
	}
	s := Semaphore{}
	s.sem = make(chan struct{}, capacity)
	return &s
}

func (s *Semaphore) Acquire() {
	s.sem <- struct{}{}
}

func (s *Semaphore) AcquireN(n int) {
	for i := 0; i < n; i++ {
		s.Acquire()
	}
}

func (s *Semaphore) ReleaseN(n int) {
	for i := 0; i < n; i++ {
		s.Release()
	}
}

func (s *Semaphore) TryAcquire() bool {
	select {
	case s.sem <- struct{}{}:
		return true
	default:
		return false
	}
}

func (s *Semaphore) Release() {
	select {
	case <-s.sem:
		return
	default:
		panic("released without acquire")
	}
}

func (s *Semaphore) Available() int {
	return cap(s.sem) - len(s.sem)
}
