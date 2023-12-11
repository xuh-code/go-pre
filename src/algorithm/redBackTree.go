package main

import (
	"fmt"
	"strings"
)

//type RBTree struct {
//	root *Node
//}

type Color bool

const (
	RED   Color = false
	BLACK Color = true
)

type Node struct {
	color  Color
	left   *Node
	right  *Node
	parent *Node
	key    int64
	value  string
}

func (n *Node) GetKey() int64 {
	return n.key
}

func (n *Node) GetValue() string {
	return n.value
}

// 父节点
func (n *Node) father() *Node {
	return n.parent
}

// 祖父节点
func (n *Node) grandParent() *Node {
	g := n.father()
	if g == nil {
		return nil
	}
	return g.parent
}

// 兄弟节点
func (n *Node) sibling() *Node {
	p := n.father()
	if p == nil {
		return nil
	}
	if n == p.left {
		return p.right
	}
	return p.left
}

// 叔叔节点
func (n *Node) uncle() *Node {
	f := n.father()
	g := n.grandParent()
	if g == nil {
		return nil
	}
	return f.sibling()
}

// 最小节点
func (n *Node) minimum() *Node {
	l := n.left
	for l != nil {
		l = l.left
	}
	return l
}

// 最大节点
func (n *Node) maximum() *Node {
	r := n.right
	for r != nil {
		r = r.right
	}
	return r
}

// 节点颜色
func (n *Node) colorOf() Color {
	if n == nil {
		return BLACK
	}
	return n.color
}

// 后继节点
func (n *Node) successor() *Node {
	if n.right != nil {
		return n.right.minimum()
	}

	p := n.parent
	c := n
	if p != nil && c == p.right {
		c = p
		p = p.parent
	}
	return p
}

// 前驱结点
func (n *Node) predecessor() *Node {
	if n.left != nil {
		return n.left.maximum()
	}

	p := n.parent
	c := n
	if p != nil && c == p.left {
		c = p
		p = p.parent
	}

	return p
}

func (n *Node) traverse(fn func(node *Node)) {
	if n == nil {
		return
	}

	n.left.traverse(fn)
	fn(n)
	n.right.traverse(fn)
}

func (n *Node) preorder() {
	//fmt.Printf("(%v %v)", n.key, n.value)
	//if n.parent == nil {
	//	fmt.Printf("nil")
	//} else {
	//	fmt.Printf("whose parent is %v", n.parent.key)
	//}
	//if n.color == RED {
	//	fmt.Println(" and color RED")
	//} else {
	//	fmt.Println(" and color BLACK")
	//}
	//if n.left != nil {
	//	fmt.Printf("%v's left child is ", n.key)
	//	n.left.preorder()
	//}
	//if n.right != nil {
	//	fmt.Printf("%v's right child is ", n.key)
	//	n.right.preorder()
	//}
}

func printTree(root *Node, indent int) {
	if root == nil {
		return
	}

	// Print right subtree
	printTree(root.right, indent+4)

	// Print current node
	fmt.Printf("%s%d(%s)\n", strings.Repeat(" ", indent), root.key, colorToString(root.color))

	// Print left subtree
	printTree(root.left, indent+4)
}

func colorToString(c Color) string {
	if c == RED {
		return "Red"
	}
	return "Black"
}

func NewNode(key int64, value string) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

type Tree struct {
	root *Node
	size int64
}

func (t *Tree) Root() *Node {
	return t.root
}

func (t *Tree) insert(item *Node) {
	var i *Node
	x := t.root

	for x != nil {
		i = x
		if item.key < x.key {
			// insert value into the left node
			x = x.left
		} else if item.key > x.key {
			// insert value into the right node
			x = x.right
		} else {
			// value exists
			return
		}
	}
	t.size++
	item.parent = i
	item.color = RED

	if i == nil {
		item.color = BLACK
		t.root = item
	} else if item.key < i.key {
		i.left = item
	} else {
		i.right = item
	}

	// Checking RBT conditions and repairing the node
	t.insertRepairNode(item)
}

// Checking RBT conditions and repairing the node
func (t *Tree) insertRepairNode(x *Node) {
	// N's parent (P) is not black
	var y *Node

	// 新插入节点, 不是根节点, 并且父节点是红色的
	for x != t.root && x.parent.color == RED {
		//	x的父节点是左节点
		if x.parent == x.grandParent().left {
			y = x.grandParent().right
			if y.colorOf() == RED {
				// Case 1: N's uncle (y) is red
				x.parent.color = BLACK
				y.color = BLACK
				x = x.grandParent()
			} else {
				// Case 2: N's uncle (y) is black, and N is a right child
				if x == x.parent.right {
					x = x.parent
					//	左旋
					t.leftRotate(x)
				}
				// Case 3: N's uncle (y) is black, and N is a left child
				x.parent.color = BLACK
				x.grandParent().color = RED
				t.rightRotate(x.grandParent())
			}
		} else {
			// Symmetric cases for the right side of the tree
			// (mirrored versions of Cases 1, 2, and 3)
			y = x.grandParent().left
			if y.colorOf() == RED {
				x.parent.color = BLACK
				y.color = BLACK
				x = x.grandParent()
			} else {
				if x == x.parent.left {
					x = x.parent
					t.rightRotate(x)
				}

				x.parent.color = BLACK
				x.grandParent().color = RED
				t.leftRotate(x.grandParent())
			}
		}
	}

	// N is the root node, i.e., first node of red–black tree
	t.root.color = BLACK
}

// 左旋
//
//			 g						g
//			  \						 \
//				p					  r1
//				 \					 /  \
//				  r1				p    r2
//	             /  \				\      \
//				l 	r2				 l      i
//					  \
//					   i
func (t *Tree) leftRotate(p *Node) {
	// Default node inserted will be a red node
	r1 := p.right
	p.right = r1.left

	if r1.left != nil {
		r1.left.parent = p
	}
	r1.parent = p.parent

	if p.parent == nil {
		t.root = r1
	} else {
		if p == p.parent.left {
			p.parent.left = r1
		} else {
			p.parent.right = r1
		}
	}
	r1.left = p
	p.parent = r1
}

// 右旋
//
//		 g			    g
//		  \			     \
//		   p			 l1
//	      / \			 / \
//	     l1			    l2  p
//	    / \                /
//	   l2  r		      r
func (t *Tree) rightRotate(p *Node) {
	l1 := p.left
	p.left = l1.right
	if l1.right != nil {
		l1.right.parent = p
	}
	l1.parent = p.parent

	if p.parent == nil {
		t.root = l1
	} else {
		if p.parent.left == p {
			p.parent.left = l1
		} else {
			p.parent.right = l1
		}
	}

	l1.right = p
	p.parent = l1
}

func (t *Tree) replace(a, b *Node) {
	if a.parent == nil {
		t.root = b
	} else if a.parent.left == a {
		a.parent.left = b
	} else {
		a.parent.right = b
	}
	if b != nil {
		b.parent = a.parent
	}
}

func (t *Tree) Search(key int64) *Node {
	pos := t.root

	if pos == nil {
		return nil
	}

	for pos != nil {
		switch {
		case pos.key == key:
			return pos
		case pos.key < key:
			pos = pos.right
		case pos.key > key:
			pos = pos.left
		}
	}
	return nil
}

// 删除
func (t *Tree) Delete(key int64) {
	z := t.Search(key)
	if z != nil {
		t.delete(z)
	}
}

// 删除
//
//	  g             g
//	   \             \
//	    p            l1
//	   / \           / \
//	  d             l2  p
//	 / \               /
//	l2  r             r
func (t *Tree) delete(p *Node) {
	t.size--

	// If strictly internal, copy successor's element to p and then make p
	if p.left != nil && p.right != nil {
		s := p.successor()
		p.key = s.key
		p.value = s.value
		p = s

		//if s.parent.left == s {
		//	s.parent.left = nil
		//}else {
		//	s.parent.right = nil
		//}
	} // p has 2 children

	// Start fixup at replacement node, if it exists.
	var replacement *Node
	if p.left != nil {
		replacement = p.left
	} else {
		replacement = p.right
	}

	if replacement != nil {
		// Link replacement to parent
		replacement.parent = p.parent
		if p.parent == nil {
			t.root = replacement
		} else if p == p.parent.left {
			p.parent.left = replacement
		} else {
			p.parent.right = replacement
		}
		// Null out links so they are OK to use by fixAfterDeletion.
		p.left = nil
		p.right = nil
		p.parent = nil

		// Fix replacement
		if p.color == BLACK {
			t.deleteRepairNode(replacement)
		}
	} else if p.parent == nil {
		// return if we are the only node.
		t.root = nil
	} else {
		//  No children. Use self as phantom replacement and unlink.
		if p.color == BLACK {
			t.deleteRepairNode(p)
		}

		if p.parent != nil {
			if p == p.parent.left {
				p.parent.left = nil
			} else if p == p.parent.right {
				p.parent.right = nil
			}
			p.parent = nil
		}
	}
}

// 删除后修复
func (t *Tree) deleteRepairNode(x *Node) {

	if x == nil {
		return
	}
	var w *Node
	for x != t.root && x.color == BLACK {
		if x == x.parent.left {
			w = x.sibling()
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.leftRotate(x.parent)
				w = x.parent.right
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.right.color == BLACK {
					w.left.color = BLACK
					w.color = RED
					t.rightRotate(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				w.right.color = BLACK
				t.leftRotate(x.parent)
				x = t.root
			}
		} else {
			w = x.sibling()
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.rightRotate(x.parent)
				w = x.parent.left
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					t.leftRotate(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				w.left.color = BLACK
				t.rightRotate(x.parent)
				x = t.root
			}

		}
	}
	x.color = BLACK
}

// NewTree returns a new rbtree
func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(key int64, value string) {
	x := NewNode(key, value)
	// Normal BST insertion
	t.insert(x)
}

// String returns a string representation of container
func (t *Tree) String() string {
	str := "RedBlackTree\n"
	if t.root != nil {
		output(t.root, "", true, &str)
	}
	return str
}

func output(node *Node, prefix string, isTail bool, str *string) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.left, newPrefix, true, str)
	}
}

func (n *Node) String() string {
	s := "R"
	if n.color == BLACK {
		s = "B"
	}
	return fmt.Sprintf("%v(%s)", n.key, s)
}

func main() {
	tree := NewTree()

	tree.Insert(1, "1")
	s := tree.String()
	fmt.Println(s)
	tree.Insert(2, "2")
	s = tree.String()
	fmt.Printf(s)
	tree.Insert(3, "3")
	s = tree.String()
	fmt.Printf(s)
	tree.Insert(4, "3")
	s = tree.String()
	fmt.Printf(s)
	tree.Insert(5, "3")
	s = tree.String()
	fmt.Printf(s)
	tree.Insert(6, "3")
	s = tree.String()
	fmt.Printf(s)
	tree.Insert(7, "3")
	s = tree.String()
	fmt.Printf(s)
	tree.Insert(8, "3")
	s = tree.String()
	fmt.Printf(s)
	tree.Insert(9, "3")
	s = tree.String()
	fmt.Printf(s)
	//tree.Insert(8, "3")
	//tree.Insert(9, "3")
	//tree.Insert(10, "3")

	//n := tree.Search(1)
	//printTree(tree.root, 0)
	//if n.value == "1" {
	//	fmt.Println("Error whilst insertion")
	//}
}
