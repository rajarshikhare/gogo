package avltree

type Node[T any] struct {
	height      int
	data        T
	left, right *Node[T]
}

type Tree[T any] struct {
	size       int
	root       *Node[T]
	comparator func(a, b T) bool
	isEqual    func(a, b T) bool
}

func (t *Tree[T]) Insert(data T) *Node[T] {
	t.root = t.insertNode(t.root, data)
	t.size++
	return t.root
}

func (t *Tree[T]) insertNode(node *Node[T], data T) *Node[T] {
	if node == nil {
		return &Node[T]{data: data, height: 1}
	}

	if t.comparator(data, node.data) {
		node.left = t.insertNode(node.left, data)
	} else {
		node.right = t.insertNode(node.right, data)
	}

	node.height = 1 + max(height(node.left), height(node.right))

	balance := getBalance(node)

	// Left Left Case
	if balance > 1 && t.comparator(data, node.left.data) {
		return rightRotate(node)
	}

	// Right Right Case
	if balance < -1 && !t.comparator(data, node.right.data) {
		return leftRotate(node)
	}

	// Left Right Case
	if balance > 1 && !t.comparator(data, node.left.data) {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// Right Left Case
	if balance < -1 && t.comparator(data, node.right.data) {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func (t *Tree[T]) Delete(data T) *Node[T] {
	t.root = t.deleteNode(t.root, data)
	if t.root != nil {
		t.size--
	}
	return t.root
}

func (t *Tree[T]) deleteNode(node *Node[T], data T) *Node[T] {
	// Implementation of delete operation
	// Similar to insert, it involves finding the node and rebalancing
	return nil // Placeholder
}

func (t *Tree[T]) Get(data T) *Node[T] {
	currentNode := t.root
	for currentNode != nil {
		if t.isEqual(currentNode.data, data) {
			return currentNode
		}
		if t.comparator(data, currentNode.data) {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}
	return nil
}

func height[T any](node *Node[T]) int {
	if node == nil {
		return 0
	}
	return node.height
}

func getBalance[T any](node *Node[T]) int {
	if node == nil {
		return 0
	}
	return height(node.left) - height(node.right)
}

// Implement rightRotate, leftRotate, etc.
