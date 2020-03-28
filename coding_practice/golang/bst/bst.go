package bst

type bstNode struct {
	key string
	val string
	left *bstNode
	right *bstNode
}

type BST struct {
	root *bstNode
	size int
}

func NewBST() *BST {
	return &BST{}
}

func (this *BST) Insert(key, value string) bool {
	if this.size == 0 {
		this.root = &bstNode{
			key: key,
			val: value,
		}
		this.size++
		return false
	}

	cur := this.root
	for {
		if cur.key == key {
			cur.val = value
			return true
		} else if key < cur.key {
			if cur.left == nil {
				cur.left = &bstNode{
					key: key,
					val: value,
				}
				this.size++
				return false
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &bstNode{
					key: key,
					val: value,
				}
				this.size++
				return false
			}
			cur = cur.right
		}
	}

	// deadcode
	return false
}

func (this *BST) InsertRecursive(key, value string) bool {
	var overwritten bool
	this.root, overwritten = this.insertRecursive(this.root, key, value)
	return overwritten
}

func (this *BST) insertRecursive(node *bstNode, key, value string) (*bstNode, bool) {
	if node == nil {
		this.size++
		return &bstNode{
			key: key,
			val: value,
		}, false
	}

	var exist bool
	if node.key == key {
		node.val = value
		exist = true
	} else if key < node.key {
		node.left, exist = this.insertRecursive(node.left, key, value)
	} else {
		node.right, exist = this.insertRecursive(node.right, key, value)
	}
	return node, exist
}

func (this *BST) Len() int {
	return this.size
}

func (this *BST) Start() *bstIterator {
	return NewBSTIterator(this.root)
}

func (this *BST) LowerBound(key string) *bstIterator {
	if this.root.key == key {
		return &bstIterator{
			curNode: this.root,
		}
	}

	ancestors := []*bstNode{}
	cur := this.root

	for cur != nil || len(ancestors) > 0 {
		for cur != nil {
			ancestors = append(ancestors, cur)
			cur = cur.left
		}
		l := len(ancestors)
		cur = ancestors[l-1]
		ancestors = ancestors[:l-1]

		if key <= cur.key {
			return &bstIterator{
				curNode: cur,
				ancestors: ancestors,
			}
		}

		cur = cur.right
	}

	return &bstIterator{
		curNode: nil,
	}
}

func (this *BST) UpperBound(key string) *bstIterator {
	iter := this.LowerBound(key)
	if iter.Done() || iter.curNode.key > key {
		return iter
	}
	iter.Next()
	return iter
}

func (this *BST) Find(key string) *bstIterator {
	iter := this.LowerBound(key)
	if !iter.Done() && iter.curNode.key == key {
		return iter
	}
	return &bstIterator{curNode: nil}
}

func (this *BST) Delete(iter *bstIterator) {
	// why do we need to delete an iterator?
}