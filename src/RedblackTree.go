package src

import (
	"cmp"
	"fmt"
)

type RedblackTree[T cmp.Ordered] struct {
	sentinel *Node[T]
	root     *Node[T]
	size     int
}

func NewRedblackTree[T cmp.Ordered]() *RedblackTree[T] {
	sentinel := &Node[T]{}

	return &RedblackTree[T]{
		sentinel: sentinel,
		root:     sentinel,
		size:     0,
	}
}

func Insert[T cmp.Ordered](data T, tree *RedblackTree[T]) {

	node := NewNode(data)
	node.color = true
	node.left = tree.sentinel
	node.right = tree.sentinel

	parent := tree.sentinel
	current := tree.root

	for current != tree.sentinel {
		parent = current

		result := cmp.Compare(node.data, current.data)
		if result == 0 {
			current.count++
			tree.size++
			return
		}

		if result < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}

	node.parent = parent
	if parent == tree.sentinel {
		tree.root = node
	} else if cmp.Compare[T](node.data, parent.data) < 0 {
		parent.left = node
	} else {
		parent.right = node
	}

	fixInsert(node, tree)
	tree.size++
}

func fixInsert[T cmp.Ordered](z *Node[T], tree *RedblackTree[T]) {
	for z.parent != tree.sentinel && z.parent.color == true {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.color == true {
				z.parent.color = false
				y.color = false
				z.parent.parent.color = true
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					rotateLeft(z, tree)
				}
				z.parent.color = false
				z.parent.parent.color = true
				rotateRight(z.parent.parent, tree)
			}
		} else {
			y := z.parent.parent.left
			if y.color == true {
				z.parent.color = false
				y.color = false
				z.parent.parent.color = true
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					rotateRight(z, tree)
				}
				z.parent.color = false
				z.parent.parent.color = true
				rotateLeft(z.parent.parent, tree)
			}
		}
	}
	tree.root.color = false
}

func rotateLeft[T cmp.Ordered](x *Node[T], tree *RedblackTree[T]) {
	y := x.right
	x.right = y.left
	if y.left != tree.sentinel {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == tree.sentinel {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func rotateRight[T cmp.Ordered](y *Node[T], tree *RedblackTree[T]) {
	x := y.left
	y.left = x.right
	if x.right != tree.sentinel {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == tree.sentinel {
		tree.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}
	x.right = y
	y.parent = x
}

func Search[T cmp.Ordered](key T, tree *RedblackTree[T]) bool {
	found := false
	node := search(key, tree.root, tree)
	if node != tree.sentinel {
		found = true
	}

	return found
}

func search[T cmp.Ordered](key T, node *Node[T], tree *RedblackTree[T]) *Node[T] {
	if node == tree.sentinel || key == node.data {
		return node
	}
	if cmp.Compare(key, node.data) < 0 {
		return search(key, node.left, tree)
	}
	return search(key, node.right, tree)
}

func minimum[T cmp.Ordered](node *Node[T], tree *RedblackTree[T]) *Node[T] {
	for node.left != tree.sentinel {
		node = node.left
	}
	return node
}

func transplant[T cmp.Ordered](u *Node[T], v *Node[T], tree *RedblackTree[T]) {
	if u.parent == tree.sentinel {
		tree.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func Delete[T cmp.Ordered](key T, tree *RedblackTree[T]) {
	z := search(key, tree.root, tree)
	if z == tree.sentinel {
		fmt.Println("Value not found in the tree:", key)
		return
	}
	if z.count > 1 {
		z.count--
		tree.size--
		return
	}
	y := z
	x := tree.sentinel
	yOriginalColor := y.color
	if z.left == tree.sentinel {
		x = z.right
		transplant(z, z.right, tree)
	} else if z.right == tree.sentinel {
		x = z.left
		transplant(z, z.left, tree)
	} else {
		y = minimum(z.right, tree)
		yOriginalColor = y.color
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			transplant(y, y.right, tree)
			y.right = z.right
			y.right.parent = y
		}
		transplant(z, y, tree)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}
	if !yOriginalColor {
		fixDelete(x, tree)
	}
	tree.size--
}

func fixDelete[T cmp.Ordered](x *Node[T], tree *RedblackTree[T]) {
	for x != tree.root && !x.color {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == true {
				w.color = false
				x.parent.color = true
				rotateLeft(x.parent, tree)
				w = x.parent.right
			}
			if !w.left.color && !w.right.color {
				w.color = true
				x = x.parent
			} else {
				if !w.right.color {
					w.left.color = false
					w.color = true
					rotateRight(w, tree)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = false
				w.right.color = false
				rotateLeft(x.parent, tree)
				x = tree.root
			}
		} else {
			w := x.parent.left
			if w.color {
				w.color = false
				x.parent.color = true
				rotateRight(x.parent, tree)
				w = x.parent.left
			}
			if !w.right.color && !w.left.color {
				w.color = true
				x = x.parent
			} else {
				if !w.left.color {
					w.right.color = false
					w.color = true
					rotateLeft(w, tree)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = false
				w.left.color = false
				rotateRight(x.parent, tree)
				x = tree.root
			}
		}
	}
	x.color = false
}

func PrintTree[T cmp.Ordered](tree *RedblackTree[T]) {
	if tree.root == tree.sentinel {
		fmt.Println("(empty tree)")
		return
	}

	printNode(tree.root, "", "", tree)
}

func printNode[T cmp.Ordered](node *Node[T], line string, indent string, tree *RedblackTree[T]) {
	if node == tree.sentinel {
		return
	}

	color := "black"
	if node.color {
		color = "red"
	}
	fmt.Printf("%s%v (%s)\n", line, node.data, color)

	rightBranch, rightIndent := "├── ", "│   "
	if node.left == tree.sentinel {
		rightBranch, rightIndent = "└── ", "    "
	}

	printNode(node.right, indent+rightBranch+"R: ", indent+rightIndent, tree)
	printNode(node.left, indent+"└── L: ", indent+"    ", tree)
}

func Size[T cmp.Ordered](tree *RedblackTree[T]) int {
	return tree.size
}

func Validate[T cmp.Ordered](tree *RedblackTree[T]) error {
	if tree.root.color {
		return fmt.Errorf("root is red")
	}
	if tree.sentinel.color {
		return fmt.Errorf("sentinel is red")
	}
	if tree.root != tree.sentinel && tree.root.parent != tree.sentinel {
		return fmt.Errorf("root's parent is not the sentinel")
	}

	var previous *Node[T]
	var walk func(node *Node[T]) (int, error)

	walk = func(node *Node[T]) (int, error) {
		if node == tree.sentinel {
			return 1, nil
		}

		leftHeight, err := walk(node.left)
		if err != nil {
			return 0, err
		}

		if previous != nil && cmp.Compare(previous.data, node.data) > 0 {
			return 0, fmt.Errorf("values are out of order at %v", node.data)
		}
		previous = node

		if node.color && (node.left.color || node.right.color) {
			return 0, fmt.Errorf("red node %v has a red child", node.data)
		}
		if node.left != tree.sentinel && node.left.parent != node {
			return 0, fmt.Errorf("bad parent link on left child of %v", node.data)
		}
		if node.right != tree.sentinel && node.right.parent != node {
			return 0, fmt.Errorf("bad parent link on right child of %v", node.data)
		}

		rightHeight, err := walk(node.right)
		if err != nil {
			return 0, err
		}
		if leftHeight != rightHeight {
			return 0, fmt.Errorf("black heights differ at %v", node.data)
		}

		if !node.color {
			leftHeight++
		}
		return leftHeight, nil
	}

	_, err := walk(tree.root)
	return err
}
