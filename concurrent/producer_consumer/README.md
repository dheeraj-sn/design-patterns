# Circular Queue with Producer-Consumer Pattern

## Overview

This implementation provides a **thread-safe circular queue** using semaphores to implement the classic **producer-consumer pattern**. Producers can enqueue items when space is available, and consumers can dequeue items when data is available.

## Data Structure

```go
type CircularQueue[T any] struct {
    capacity    int64           // Maximum number of elements
    front       int64           // Index of the front element (oldest)
    back        int64           // Index where next element will be inserted
    data        []T             // Underlying array buffer
    currentSize int64           // Current number of elements
    lock        sync.Mutex      // Protects shared state modifications
    emptySlots  *semaphore.Weighted  // Tracks available empty slots
    fullSlots   *semaphore.Weighted  // Tracks available full slots
}
```

## Semaphore-Based Synchronization

The queue uses **two semaphores** to coordinate producers and consumers:

### Initialization
- **`emptySlots`**: Initialized with `capacity` permits (all slots are empty initially)
- **`fullSlots`**: All `capacity` permits are **pre-acquired** (no full slots initially)

This ensures:
- **Producers** must wait for an empty slot before enqueueing
- **Consumers** must wait for a full slot before dequeueing

## Core Operations

### Enqueue (Producer)

```go
func Enqueue(value T) error
```

**Logic Flow:**
1. **Acquire empty slot**: `emptySlots.Acquire(1)` - Blocks if queue is full
2. **Lock mutex**: Protect critical section
3. **Insert element**: Store value at `data[back]`
4. **Update indices**: 
   - Increment `currentSize`
   - Advance `back` using circular indexing
5. **Release full slot**: `fullSlots.Release(1)` - Signal that a slot is now full
6. **Unlock mutex**

**Key Point**: After acquiring `emptySlots`, the queue **cannot be full**, so the operation always succeeds.

### Dequeue (Consumer)

```go
func Dequeue() (bool, T)
```

**Logic Flow:**
1. **Acquire full slot**: `fullSlots.Acquire(1)` - Blocks if queue is empty
2. **Lock mutex**: Protect critical section
3. **Remove element**: Read value from `data[front]`
4. **Update indices**:
   - Decrement `currentSize`
   - Advance `front` using circular indexing
5. **Release empty slot**: `emptySlots.Release(1)` - Signal that a slot is now empty
6. **Unlock mutex**

**Key Point**: After acquiring `fullSlots`, the queue **cannot be empty**, so the operation always succeeds.

## Circular Indexing

The queue uses **modular arithmetic** to wrap indices around:

```go
func GetIndex(currentIndex int64, stepSize int64) int64 {
    return (currentIndex + stepSize + capacity) % capacity
}
```

**Example** (capacity = 3):
- Index 0 → step +1 → Index 1
- Index 2 → step +1 → Index 0 (wraps around)
- Index 0 → step -1 → Index 2 (wraps backward)

The `+ capacity` ensures the result is always positive before modulo operation, handling negative step sizes correctly.

## Thread Safety

1. **Semaphores**: Control blocking behavior (producers wait for space, consumers wait for data)
2. **Mutex**: Protects the critical section where shared state (`front`, `back`, `currentSize`, `data`) is modified
3. **Order**: Semaphore acquire happens **before** mutex lock to avoid deadlocks

## Producer-Consumer Flow

```
Initial State:
- emptySlots: capacity permits available
- fullSlots: 0 permits available (all pre-acquired)

Producer (Enqueue):
  emptySlots.Acquire() → Lock → Insert → fullSlots.Release() → Unlock

Consumer (Dequeue):
  fullSlots.Acquire() → Lock → Remove → emptySlots.Release() → Unlock
```

This ensures:
- ✅ Producers block when queue is full
- ✅ Consumers block when queue is empty
- ✅ No race conditions on shared state
- ✅ Proper synchronization between multiple producers/consumers

## Helper Methods

- **`GetFront()`**: Returns the front element without removing it (thread-safe)
- **`GetBack()`**: Returns the last inserted element (thread-safe)
- **`IsEmpty()` / `IsFull()`**: Thread-safe state queries

