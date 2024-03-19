package util

import (
	"sync"
)

// SafeMap 是一个线程安全的泛型映射结构体。
// 它使用读写锁（sync.RWMutex）来保证在并发环境下的安全访问。
// K 为映射的键类型，V 为映射的值类型，K需要是可比较的类型。
type SafeMap[K comparable, V any] struct {
	dataLock sync.RWMutex // 用于保护 dataMap 的读写锁
	dataMap  map[K]V      // 存储实际映射数据的内部 map
}

// NewSafeMap 创建并返回一个新的 SafeMap 实例。
// 返回的 SafeMap 实例是空的，准备好接受键值对。
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		dataLock: sync.RWMutex{},
		dataMap:  make(map[K]V),
	}
}

// Set 在 SafeMap 中设置一个键值对。
// 如果键已存在，其对应的值将被新值覆盖。
// 此操作是线程安全的。
func (sm *SafeMap[K, V]) Set(key K, value V) {
	sm.dataLock.Lock()      // 获取写锁
	sm.dataMap[key] = value // 设置键值对
	sm.dataLock.Unlock()    // 释放写锁
}

// Get 从 SafeMap 中获取一个值，以及该值是否存在的标志。
// 如果键存在于映射中，则返回该键对应的值和 true；否则，返回值类型的零值和 false。
// 此操作是线程安全的。
func (sm *SafeMap[K, V]) Get(key K) (V, bool) {
	sm.dataLock.RLock()              // 获取读锁
	value, exists := sm.dataMap[key] // 尝试获取键对应的值
	sm.dataLock.RUnlock()            // 释放读锁
	return value, exists
}

// Delete 从 SafeMap 中删除一个键及其对应的值。
// 如果键存在，则执行删除操作；否则，不执行任何操作。
// 此操作是线程安全的。
func (sm *SafeMap[K, V]) Delete(key K) {
	sm.dataLock.Lock()      // 获取写锁
	delete(sm.dataMap, key) // 删除键值对
	sm.dataLock.Unlock()    // 释放写锁
}

// ReplaceAll 用一个新的 map 替换 SafeMap 中的所有数据。
// 这个操作会先清除所有现有的数据，然后设置新的数据。
// 此操作是线程安全的。
func (sm *SafeMap[K, V]) ReplaceAll(newMap map[K]V) {
	sm.dataLock.Lock()   // 获取写锁
	sm.dataMap = newMap  // 替换整个 map
	sm.dataLock.Unlock() // 释放写锁
}

// Range 对 SafeMap 中的每一个键值对执行一个回调函数。
// 如果回调函数对某个键值对返回 false，则立即停止遍历。
// 此操作在执行期间会阻塞写操作，但允许其他读操作并行执行。
// 注意：回调函数中不应进行任何修改 SafeMap 的操作，以避免死锁。
func (sm *SafeMap[K, V]) Range(f func(key K, value V) bool) {
	sm.dataLock.RLock()         // 获取读锁
	defer sm.dataLock.RUnlock() // 确保函数结束时释放读锁
	for k, v := range sm.dataMap {
		if !f(k, v) { // 执行回调，如果返回 false，则终止遍历
			break
		}
	}
}

// Len 返回 SafeMap 中的键值对数量。
// 此操作是线程安全的。
func (sm *SafeMap[K, V]) Len() int {
	sm.dataLock.RLock()         // 获取读锁
	defer sm.dataLock.RUnlock() // 确保函数结束时释放读锁
	return len(sm.dataMap)      // 返回内部 map 的长度

}
