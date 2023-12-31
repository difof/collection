package collection

// Map is a generic container for map.
type Map[K comparable, V any] struct {
	m map[K]V
}

// NewMap creates a new Map.
func NewMap[K comparable, V any](items ...Tuple[K, V]) *Map[K, V] {
	var m map[K]V

	if items != nil && len(items) > 0 {
		m = make(map[K]V, len(items))
		for _, item := range items {
			m[item.Key()] = item.Value()
		}
	} else {
		m = make(map[K]V)
	}

	return &Map[K, V]{m: m}
}

func AsMap[K comparable, V any](c Collection[K, V, Tuple[K, V]]) *Map[K, V] {
	return c.(*Map[K, V])
}

// Get gets a value from the map.
func (m *Map[K, V]) Get(key K) (v V, ok bool) {
	v, ok = m.m[key]
	return
}

// Set sets a value in the map.
func (m *Map[K, V]) Set(key K, val V) {
	m.m[key] = val
}

// SetElem sets a value in the map.
func (m *Map[K, V]) SetElem(elem Tuple[K, V]) {
	m.Set(elem.Key(), elem.Value())
}

// AppendElem appends an element to the map.
func (m *Map[K, V]) AppendElem(elem Tuple[K, V]) {
	m.SetElem(elem)
}

// Delete deletes a value from the map.
func (m *Map[K, V]) Delete(key K) {
	delete(m.m, key)
}

// Len returns the length of the map.
func (m *Map[K, V]) Len() int {
	return len(m.m)
}

// Cap returns the capacity of the map, which is equal to the length.
func (m *Map[K, V]) Cap() int {
	return m.Len()
}

// Values returns all values in the map.
func (m *Map[K, V]) Values() []V {
	v := make([]V, 0, m.Len())
	for _, val := range m.m {
		v = append(v, val)
	}

	return v
}

// Keys returns all keys in the map.
func (m *Map[K, V]) Keys() []K {
	k := make([]K, 0, m.Len())
	for key := range m.m {
		k = append(k, key)
	}

	return k
}

// Clear clears the map.
func (m *Map[K, V]) Clear() {
	m.m = make(map[K]V)
}

// HasKey checks if the map has the key.
func (m *Map[K, V]) HasKey(key K) bool {
	_, ok := m.m[key]
	return ok
}

// IsEmpty checks if the map is empty.
func (m *Map[K, V]) IsEmpty() bool {
	return len(m.m) == 0
}

func (m *Map[K, V]) Iter() *Iterator[Tuple[K, V]] {
	return NewIterator(m.AsIterable())
}

func (m *Map[K, V]) IterHandler(iter *Iterator[Tuple[K, V]]) {
	go func() {
		for k, v := range m.m {
			select {
			case <-iter.Done():
				return
			case iter.NextChannel() <- NewTuple(k, v):
			}
		}

		iter.IterationDone()
	}()
}

// Clone returns a copy of the map.
func (m *Map[K, V]) Clone() Collection[K, V, Tuple[K, V]] {
	items := make([]Tuple[K, V], 0, m.Len())
	for k, v := range m.m {
		items = append(items, NewTuple(k, v))
	}

	return NewMap[K, V](items...)
}

// Factory returns a new instance of the map.
func (m *Map[K, V]) Factory() Collection[K, V, Tuple[K, V]] {
	return NewMap[K, V]()
}

func (m *Map[K, V]) FactoryFrom(values []V) Collection[K, V, Tuple[K, V]] {
	panic("Map doesn't support FactoryFrom")
}

func (m *Map[K, V]) AsCollection() Collection[K, V, Tuple[K, V]] {
	return m
}

func (m *Map[K, V]) AsIterable() Iterable[Tuple[K, V]] {
	return m
}
