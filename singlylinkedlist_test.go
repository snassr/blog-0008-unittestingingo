package singlylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSinglyLinkedList(t *testing.T) {
	for title, fn := range map[string]func(t *testing.T){
		"append item to list":  testAppend, // test: append functionality
		"return items in list": testValues, // test: values functionality
	} {
		t.Run(title, func(t *testing.T) {
			fn(t)
		})
	}
}

func testAppend(t *testing.T) {
	list := NewSinglyLinkedList()

	tableTests := []struct {
		list     SinglyLinkedList
		val      int
		expected []int
	}{
		{list: list, val: 1, expected: []int{1}},    // case: empty list
		{list: list, val: 2, expected: []int{1, 2}}, // case: non-empty list
	}

	for _, tt := range tableTests {
		tt.list.Append(tt.val)
		require.Equal(t, tt.expected, tt.list.Values())
	}
}

func testValues(t *testing.T) {
	tableTests := []struct {
		list     SinglyLinkedList
		expected []int
	}{
		{list: NewSinglyLinkedList(), expected: []int{}},         // case: empty list
		{list: NewSinglyLinkedList(1), expected: []int{1}},       // case: one-item list
		{list: NewSinglyLinkedList(1, 2), expected: []int{1, 2}}, // case: multiple item list
	}

	for _, tt := range tableTests {
		require.Equal(t, tt.expected, tt.list.Values())
	}
}
