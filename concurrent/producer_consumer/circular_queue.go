package producerconsumer

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

type CircularQueue[T any] struct {
	capacity    int64
	front       int64
	back        int64
	data        []T
	currentSize int64
	lock        sync.Mutex
	emptySlots  *semaphore.Weighted
	fullSlots   *semaphore.Weighted
}

func NewCircularQueue[T any](capacity int64) *CircularQueue[T] {
	c := &CircularQueue[T]{
		capacity:    capacity,
		front:       0,
		back:        0,
		currentSize: 0,
		data:        make([]T, capacity),
		emptySlots:  semaphore.NewWeighted(capacity),
		fullSlots:   semaphore.NewWeighted(capacity),
		lock:        sync.Mutex{},
	}
	c.fullSlots.Acquire(context.Background(), capacity)
	return c
}

func (s *CircularQueue[T]) isFull() bool {
	return s.currentSize == s.capacity
}

func (s *CircularQueue[T]) IsFull() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.isFull()
}

func (s *CircularQueue[T]) isEmpty() bool {
	return s.currentSize == 0
}

func (s *CircularQueue[T]) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.isEmpty()
}

func (s *CircularQueue[T]) Enqueue(value T) error {
	err := s.emptySlots.Acquire(context.Background(), 1)
	if err != nil {
		return err
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	s.currentSize += 1
	s.data[s.back] = value
	s.back = s.GetIndex(s.back, 1)
	s.fullSlots.Release(1)
	return nil
}

func (s *CircularQueue[T]) Dequeue() (bool, T) {
	s.fullSlots.Acquire(context.Background(), 1)
	s.lock.Lock()
	defer s.lock.Unlock()
	s.currentSize -= 1
	frontValue := s.data[s.front]
	s.front = s.GetIndex(s.front, 1)
	s.emptySlots.Release(1)
	return true, frontValue
}

func (s *CircularQueue[T]) GetFront() (T, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	var zero T
	if s.isEmpty() {
		return zero, false
	}
	return s.data[s.front], true
}

func (s *CircularQueue[T]) GetBack() (T, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	var zero T
	if s.isEmpty() {
		return zero, false
	}
	index := s.GetIndex(s.back, -1)
	return s.data[index], true
}

func (s *CircularQueue[T]) GetIndex(currentIndex int64, stepSize int64) int64 {
	return (currentIndex + stepSize + s.capacity) % s.capacity
}
