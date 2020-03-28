package bst

type bstIterator struct {
	curNode *bstNode // nil when there are no more nodes to iterate
	ancestors []*bstNode
}

func NewBSTIterator(node *bstNode) *bstIterator {
	ancestors := []*bstNode{}
	cur := node

	for cur != nil {
		ancestors = append(ancestors, cur)
		cur = cur.left
	}
	l := len(ancestors)
	cur = ancestors[l-1]
	ancestors = ancestors[:l-1]
	
	return &bstIterator{
		curNode: cur,
		ancestors: ancestors,
	}
}

func (this *bstIterator) Done() bool {
	return this.curNode == nil && len(this.ancestors) == 0
}

func (this *bstIterator) Value() string {
	if this.Done() {
		return ""
	}
	return this.curNode.val
}

func (this *bstIterator) Next() {
	if this.Done() {
		return
	}

	cur := this.curNode
	cur = cur.right
	for cur != nil {
		this.ancestors = append(this.ancestors, cur)
		cur = cur.left
	}
	l := len(this.ancestors)
	if l == 0 {
		this.curNode = nil
	} else {
		this.curNode = this.ancestors[l-1]
		this.ancestors = this.ancestors[:l-1]
	}
}