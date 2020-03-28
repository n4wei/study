package stack

// 4:15-4:27

type StringStack struct {
	store []string
}

func NewStringStack() *StringStack {
	// rely on append to instantiate the underlying []string and resize it when more strings are pushed
	return new(StringStack)
}

// runtime amoritized O(1)
// space O(n) for n number of strings in stack
func (this *StringStack) Push(str string) {
	this.store = append(this.store, str)
}

// runtime O(1)
func (this *StringStack) Pop() string {
	if len(this.store) == 0 {
		return ""
	}

	l := len(this.store)
	result := this.store[l-1]
	this.store = this.store[:l-1]
	return result
}

// runtime O(1)
func (this *StringStack) Len() int {
	return len(this.store)
}