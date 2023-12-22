package collection

import (
	"testing"
)

func TestForwardLinkedList_Iter(t *testing.T) {
	list := NewForwardLinkedList[int]()
	list.Append(1)
	list.Append(4)
	list.Append(2)
	list.Append(3)
	list.Append(5)

	list = AsForwardLinkedList(OrderBy(list.AsCollection(), NumericComparator[int]))

	Each(list.AsIterable(), func(i Tuple[int, int]) error {
		t.Log(i.Value())
		return nil
	})
}
