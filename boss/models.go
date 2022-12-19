package boss

import "sync"

type MessageCapacity struct {
	capacity int
	current  int
	Mutex    sync.Mutex
}

func NewMessageCapacity(capacity int) MessageCapacity {
	return MessageCapacity{
		capacity: capacity,
		current:  capacity,
	}
}

func (m *MessageCapacity) IsOverflowed() bool {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.current--

	if m.current < 0 {
		return true
	}

	return false
}

func (m *MessageCapacity) Refresh() {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.current = m.capacity
}
