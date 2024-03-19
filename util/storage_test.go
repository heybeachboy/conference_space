package util

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestNewSafeMap(t *testing.T) {
	sm := NewSafeMap[string, string]()
	if sm == nil {
		t.Errorf("NewSafeMap() = nil, want non-nil")
	}
}

func TestSetAndGet(t *testing.T) {
	sm := NewSafeMap[string, string]()
	sm.Set("key1", "value1")
	value, exists := sm.Get("key1")

	if !exists || value != "value1" {
		t.Errorf("Get() = %v, %t; want %v, %t", value, exists, "value1", true)
	}
}

func TestDelete(t *testing.T) {
	sm := NewSafeMap[string, string]()
	sm.Set("key1", "value1")
	sm.Delete("key1")
	_, exists := sm.Get("key1")

	if exists {
		t.Errorf("Delete() failed, key %v still exists", "key1")
	}
}

func TestReplaceAll(t *testing.T) {
	sm := NewSafeMap[string, string]()
	sm.Set("key1", "value1")
	newMap := map[string]string{"key2": "value2"}
	sm.ReplaceAll(newMap)

	_, exists := sm.Get("key1")
	if exists {
		t.Errorf("ReplaceAll() failed, old key %v still exists", "key1")
	}

	value, exists := sm.Get("key2")
	if !exists || value != "value2" {
		t.Errorf("ReplaceAll() failed, new key %v not set correctly", "key2")
	}
}

func TestRange(t *testing.T) {
	sm := NewSafeMap[string, string]()
	sm.Set("key1", "value1")
	sm.Set("key2", "value2")

	count := 0
	sm.Range(func(key string, value string) bool {
		count++
		return true
	})

	if count != 2 {
		t.Errorf("Range() visited %v items, want %v", count, 2)
	}
}

func TestLen(t *testing.T) {
	sm := NewSafeMap[string, string]()
	sm.Set("key1", "value1")
	sm.Set("key2", "value2")
	if sm.Len() != 2 {
		t.Errorf("Len() = %v, want %v", sm.Len(), 2)
	}
}

func TestNestedSafeMap(t *testing.T) {
	outerMap := NewSafeMap[string, *SafeMap[string, string]]()

	innerMap := NewSafeMap[string, string]()
	innerMap.Set("innerKey", "innerValue")

	outerMap.Set("outerKey", innerMap)

	retrievedInnerMap, exists := outerMap.Get("outerKey")
	if !exists {
		t.Fatal("Outer key 'outerKey' not found")
	}

	value, exists := retrievedInnerMap.Get("innerKey")
	if !exists || value != "innerValue" {
		t.Errorf("Inner map value for 'innerKey' = %v, %t; want %v, %t", value, exists, "innerValue", true)
	}
}

func TestConcurrentNestedSafeMapReadWriteWithChan(t *testing.T) {
	runTimes := int(10000)
	writeGoRuntimeNum := int(2)
	readGoRuntimeNum := int(128)
	randomCheckTimes := int(200)

	outerMap := NewSafeMap[string, *SafeMap[string, string]]()

	outerKey := "outerMapKey"
	innerMap := NewSafeMap[string, string]()
	outerMap.Set(outerKey, innerMap)

	var wg sync.WaitGroup

	keyChan := make(chan string, runTimes*writeGoRuntimeNum)

	for i := 0; i < writeGoRuntimeNum; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < runTimes; j++ {
				innerKey := fmt.Sprintf("innerKey%d-%d", i, j)
				innerValue := fmt.Sprintf("innerValue%d-%d", i, j)
				innerMap.Set(innerKey, innerValue)

				keyChan <- innerKey
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(keyChan)
	}()

	for i := 0; i < readGoRuntimeNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < runTimes; j++ {
				_, exists := outerMap.Get(outerKey)
				if !exists {
					t.Errorf("Outer key '%s' not found", outerKey)
					return
				}
			}
		}()
	}

	allKeys := make([]string, 0, runTimes*writeGoRuntimeNum)
	for key := range keyChan {
		allKeys = append(allKeys, key)
	}

	for i := 0; i < randomCheckTimes; i++ {
		index := rand.Intn(len(allKeys))
		sampleKey := allKeys[index]

		_, exists := innerMap.Get(sampleKey)
		if !exists {
			t.Errorf("Expected inner key '%s' not found in random sampling", sampleKey)
		}
	}
}

func BenchmarkSafeMapConcurrentReadWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sm := NewSafeMap[string, *SafeMap[string, string]]()

		var wg sync.WaitGroup
		// 写线程
		for w := 0; w < 2; w++ {
			wg.Add(1)
			go func(w int) {
				defer wg.Done()
				innerMap := NewSafeMap[string, string]()
				innerKey := fmt.Sprintf("innerKey%d", w)
				innerMap.Set(innerKey, "value")
				sm.Set("innerMap", innerMap)
			}(w)
		}

		// 读线程
		for r := 0; r < 128; r++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if innerMap, ok := sm.Get("innerMap"); ok {
					_, _ = innerMap.Get("innerKey1")
				}
			}()
		}

		wg.Wait()
	}
}

func BenchmarkSyncMapConcurrentReadWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sm sync.Map

		var wg sync.WaitGroup
		// 写线程
		for w := 0; w < 2; w++ {
			wg.Add(1)
			go func(w int) {
				defer wg.Done()
				innerMap := &sync.Map{}
				innerKey := fmt.Sprintf("innerKey%d", w)
				innerMap.Store(innerKey, "value")
				sm.Store("innerMap", innerMap)
			}(w)
		}

		// 读线程
		for r := 0; r < 128; r++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if innerMap, ok := sm.Load("innerMap"); ok {
					innerMap.(*sync.Map).Load("innerKey1")
				}
			}()
		}

		wg.Wait()
	}
}
