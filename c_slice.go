package collection

// Slice is a generic slice.
//
// It doesn't check out of bounds access.
type Slice[V any] struct {
	slice []V
}

// NewSlice creates a new Slice.
func NewSlice[V any](items ...V) *Slice[V] {
	if items != nil {
		return &Slice[V]{slice: items}
	}

	return &Slice[V]{slice: make([]V, 0)}
}

// NewSliceN creates a new Slice with size and capacity of n.
func NewSliceN[V any](s, n int) *Slice[V] {
	return &Slice[V]{slice: make([]V, s, n)}
}

func AsSlice[V any](c Collection[int, V, Tuple[int, V]]) *Slice[V] {
	return c.(*Slice[V])
}

// Slice returns a new Slice with the elements from start to end-1.
func (s *Slice[V]) Slice(start, end int) *Slice[V] {
	return NewSlice(s.slice[start:end]...)
}

// Append appends a value to the slice.
func (s *Slice[V]) Append(item V) {
	s.slice = append(s.slice, item)
}

// AppendElem appends a value to the slice.
func (s *Slice[V]) AppendElem(elem Tuple[int, V]) {
	s.Append(elem.Value())
}

// AppendSlice appends a slice to the slice.
func (s *Slice[V]) AppendSlice(slice []V) {
	s.slice = append(s.slice, slice...)
}

// AppendSafeSlice appends a Slice to the slice.
func (s *Slice[V]) AppendSafeSlice(other *Slice[V]) {
	s.slice = append(s.slice, other.slice...)
}

// Get gets a value from the slice.
func (s *Slice[V]) Get(i int) (V, bool) {
	return s.slice[i], true
}

// Set sets a value in the slice.
func (s *Slice[V]) Set(i int, item V) {
	s.slice[i] = item
}

// SetElem sets a value in the slice.
func (s *Slice[V]) SetElem(item Tuple[int, V]) {
	s.Set(item.Key(), item.Value())
}

// Delete deletes a value at given index from the slice.
func (s *Slice[V]) Delete(i int) {
	s.slice = append(s.slice[:i], s.slice[i+1:]...)
}

// Len returns the length of the slice.
func (s *Slice[V]) Len() int {
	return len(s.slice)
}

// Cap returns the capacity of the slice.
func (s *Slice[V]) Cap() int {
	return cap(s.slice)
}

// IsEmpty returns true if the slice is empty.
func (s *Slice[V]) IsEmpty() bool {
	return len(s.slice) == 0
}

func (s *Slice[V]) Iter() *Iterator[Tuple[int, V]] {
	return NewIterator(s.AsIterable())
}

func (s *Slice[V]) IterHandler(iter *Iterator[Tuple[int, V]]) {
	go func() {
		for i, item := range s.slice {
			select {
			case <-iter.Done():
				return
			case iter.NextChannel() <- NewTuple(i, item):
			}
		}

		iter.IterationDone()
	}()
}

// Clear clears the slice.
func (s *Slice[V]) Clear() {
	s.slice = make([]V, 0)
}

// Clone returns a clone of the slice. Same as Values.
func (s *Slice[V]) Clone() Collection[int, V, Tuple[int, V]] {
	newSlice := make([]V, len(s.slice))

	for i, item := range s.slice {
		newSlice[i] = item
	}

	return NewSlice(newSlice...)
}

// Values returns all values in the slice.
// This function is not thread-safe, use with caution.
func (s *Slice[V]) Values() []V {
	return s.slice
}

// Compare compares two slice items.
func (s *Slice[V]) Compare(i, j Tuple[int, V], comp func(V, V) CompareResult) CompareResult {
	v1, _ := s.Get(i.Key())
	v2, _ := s.Get(j.Key())
	return comp(v1, v2)
}

// Factory returns a new Slice.
func (s *Slice[V]) Factory() Collection[int, V, Tuple[int, V]] {
	return NewSlice[V]()
}

func (s *Slice[V]) FactoryFrom(values []V) Collection[int, V, Tuple[int, V]] {
	return NewSlice[V](values...)
}

func (s *Slice[V]) AsCollection() Collection[int, V, Tuple[int, V]] {
	return s
}

func (s *Slice[V]) AsIterable() Iterable[Tuple[int, V]] {
	return s
}
