package main

func main() {

}

type node struct {
	key         int
	value       int
	left, right *node
	N           int // 以该结点为根的子树中的结点数
}

func newNode(key, val, N int) *node {
	return &node{
		key:   key,
		value: val,
		N:     N,
	}
}

func (x *node) get(key int) *node {
	if x == nil {
		return nil
	}
	if key < x.key {
		x = x.left
		return x.get(key)
	} else if key > x.key {
		x = x.right
		return x.get(key)
	} else {
		return x
	}
}

func (x *node) put(key, value int) *node {
	if x == nil {
		return &node{
			key:   key,
			value: value,
			N:     1,
		}
	}
	if x.key > key {
		x.left.put(key, value)
	} else if x.key < key {
		x.right.put(key, value)
	} else {
		x.value = value
	}
	x.N = size(x.left) + size(x.right) + 1
	return x
}

func (x *node) getI(key int) *node {
	for x != nil {
		if x.key < key {
			x = x.right
		} else if x.key > key {
			x = x.left
		} else {
			return x
		}
	}
	return nil
}

func (x *node) putI(key, value int) {
	for x != nil {
		if x.key == key {
			x.value = value
			return
		} else if x.key > key {
			if x.left != nil {
				x.N++
				x = x.left
			} else {
				x.left = &node{
					key:   key,
					value: value,
					N:     0,
				}
				break
			}
		} else if x.key < key {
			if x.right != nil {
				x.N++
				x = x.right
			} else {
				x.right = &node{
					key:   key,
					value: value,
					N:     0,
				}
				break
			}
		}
	}
}

func min(x *node) *node {
	if x.left == nil {
		return x
	}
	return min(x.left)
}

func max(x *node) *node {
	if x.right == nil {
		return x
	}
	return max(x.right)
}

func deleteMin(x *node) *node {
	if x.left == nil {
		return x.right
	}
	x.left = deleteMin(x.left)
	x.N = size(x.left) + size(x.right) + 1
	return x
}

func (x *node) deleteNode(key int) *node {
	if x == nil {
		return nil
	}
	if key < x.key {
		x.left = x.left.deleteNode(key)
	} else if key > x.key {
		x.right = x.right.deleteNode(key)
	} else {
		if x.right == nil {
			return x.left
		}
		if x.left == nil {
			return x.right
		}
		t := x
		x = min(t.right)
		deleteMin(x.right)
		x.left = t.left
	}
	x.N = size(x.left) + size(x.right) + 1
	return x
}

func size(x *node) int {
	if x == nil {
		return 0
	} else {
		return x.N
	}
}
