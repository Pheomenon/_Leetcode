package main

import "testing"

func TestInsert(t *testing.T) {
	root := newNode(21, 21, 0)
	root.putI(10, 10)
	root.putI(6, 6)
	root.putI(34, 34)
	root.putI(65, 65)
	root.putI(1, 1)
	root.putI(3, 3)
	root.putI(54, 54)
	root.putI(28, 28)
	root.putI(43, 43)

	val := root.get(28)
	if val.value != 28 {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	root := newNode(21, 21, 0)
	root.putI(10, 10)
	root.putI(6, 6)
	root.putI(34, 34)
	root.putI(65, 65)
	root.putI(1, 1)
	root.putI(3, 3)
	root.putI(54, 54)
	root.putI(28, 28)
	root.putI(43, 43)

	val := root.deleteNode(54)
	if val == root.get(1) {
		t.Fail()
	}
}
