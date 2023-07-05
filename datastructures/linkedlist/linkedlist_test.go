package linkedlist_test

import (
	"testing"

	"github.com/godarkthrt/godsalgo/datastructures/linkedlist"
)

func TestPrintLinkedList(t *testing.T) {
	ll := linkedlist.NewLinkedList(10)
	ll.PrintList()
}
