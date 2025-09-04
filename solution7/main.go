package main

import (
	"errors"
	"fmt"
	"sync"
)

type SafeMap struct {
	mu   sync.Mutex
	data map[string]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]interface{}),
	}
}

func (sm *SafeMap) Add(key string, value interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (interface{}, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if value, ok := sm.data[key]; ok == true {
		return value, nil
	}
	err := errors.New("Key does not exist")
	return "error", err
}

func main() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", idx)
			value := fmt.Sprintf("value%d", idx)
			safeMap.Add(key, value)
		}(i)
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", idx)
			if value, err := safeMap.Get(key); err == nil {
				fmt.Printf("Receive: %s = %s\n", key, value)
			}
		}(i)
	}
	wg.Wait()
}
