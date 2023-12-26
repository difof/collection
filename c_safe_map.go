package collection

import "sync"

type SafeMap[K comparable, V any] struct {
	m    Map[K, V]
	lock sync.Mutex
}

// NewSafeMap creates a new SafeMap.
func NewSafeMap[K comparable, V any](items ...Tuple[K, V]) *SafeMap[K, V] {
	return &SafeMap[K, V]{m: *NewMap[K, V](items...)}
}

func AsSafeMap[K comparable, V any](c Collection[K, V, Tuple[K, V]]) *SafeMap[K, V] {
	return c.(*SafeMap[K, V])
}

func (m *SafeMap[K, V]) Get(key K) (V, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.m.Get(key)
}

func (m *SafeMap[K, V]) Set(key K, val V) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.m.Set(key, val)
}

func (m *SafeMap[K, V]) SetElem(elem Tuple[K, V]) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.m.SetElem(elem)
}

func (m *SafeMap[K, V]) AppendElem(elem Tuple[K, V]) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.m.AppendElem(elem)
}

func (m *SafeMap[K, V]) Delete(key K) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.m.Delete(key)
}

func (m *SafeMap[K, V]) Len() int {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.m.Len()
}

func (m *SafeMap[K, V]) Cap() int {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.m.Cap()
}

func (m *SafeMap[K, V]) Values() []V {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.m.Values()
}

func (m *SafeMap[K, V]) Keys() []K {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.m.Keys()
}

func (m *SafeMap[K, V]) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.m.Clear()
}

func (m *SafeMap[K, V]) HasKey(key K) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.m.HasKey(key)
}

func (m *SafeMap[K, V]) IsEmpty() bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.m.IsEmpty()
}

func (m *SafeMap[K, V]) Iter() *Iterator[Tuple[K, V]] {
	return m.m.Iter()
}

func (m *SafeMap[K, V]) IterHandler(iter *Iterator[Tuple[K, V]]) {
	go func() {
		m.lock.Lock()
		defer m.lock.Unlock()

		for k, v := range m.m.m {
			select {
			case <-iter.Done():
				return
			case iter.NextChannel() <- NewTuple(k, v):
			}
		}

		iter.IterationDone()
	}()
}

func (m *SafeMap[K, V]) Clone() Collection[K, V, Tuple[K, V]] {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.m.Clone()
}

func (m *SafeMap[K, V]) Factory() Collection[K, V, Tuple[K, V]] {
	return NewSafeMap[K, V]()
}

func (m *SafeMap[K, V]) FactoryFrom(values []V) Collection[K, V, Tuple[K, V]] {
	panic("SafeMap doesn't support FactoryFrom")
}

func (m *SafeMap[K, V]) AsCollection() Collection[K, V, Tuple[K, V]] {
	return m
}

func (m *SafeMap[K, V]) AsIterable() Iterable[Tuple[K, V]] {
	return m
}
