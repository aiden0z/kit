package list

type Element struct {
	next, prev *Element
	list       *DoublyLinkedList
	Value      interface{}
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// double link ring list
type DoublyLinkedList struct {
	root Element
	size int
}

func (l *DoublyLinkedList) Init() *DoublyLinkedList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.size = 0
	return l
}

// create an initialized list.
func NewDoublyLinkedList() *DoublyLinkedList { return new(DoublyLinkedList).Init() }

func (l *DoublyLinkedList) Len() int { return l.size }

// retrun the fisrt element of list or nil
func (l *DoublyLinkedList) First() *Element {
	if l.size == 0 {
		return nil
	}
	return l.root.next
}

func (l *DoublyLinkedList) Last() *Element {
	if l.size == 0 {
		return nil
	}
	return l.root.prev
}

func (l *DoublyLinkedList) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.size++
	return e
}

func (l *DoublyLinkedList) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

func (l *DoublyLinkedList) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.size--
	return e
}

func (l *DoublyLinkedList) Remove(e *Element) interface{} {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

func (l *DoublyLinkedList) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *DoublyLinkedList) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *DoublyLinkedList) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

func (l *DoublyLinkedList) InsertBefore(v interface{}, at *Element) *Element {
	if at.list != l {
		return nil
	}
	return l.insertValue(v, at.prev)
}

func (l *DoublyLinkedList) InsertAfter(v interface{}, at *Element) *Element {
	if at.list != l {
		return nil
	}
	return l.insertValue(v, at)
}

func (l *DoublyLinkedList) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.insert(l.remove(e), &l.root)
}

func (l *DoublyLinkedList) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.insert(l.remove(e), l.root.prev)
}

// move element e to its new position before at
func (l *DoublyLinkedList) MoveBefore(e, at *Element) {
	if e.list != l || e == at || at.list != l {
		return
	}
	l.insert(l.remove(e), at.prev)

}

func (l *DoublyLinkedList) MoveAfter(e, at *Element) {
	if e.list != l || e == at || at.list != l {
		return
	}
	l.insert(l.remove(e), at)
}

// inserts a copy of an other list at the back of list l
func (l *DoublyLinkedList) PushBackList(other *DoublyLinkedList) {
	l.lazyInit()
	for i, e := other.Len(), other.First(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// inserts a copy of other list at the front of list l
func (l *DoublyLinkedList) PushFrontList(other *DoublyLinkedList) {
	l.lazyInit()
	for i, e := other.Len(), other.Last(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

type Item struct {
	Value interface{}
	next  *Item
	list  *List
}

func (e *Item) Next() *Item {
	if p := e.next; e.list != nil {
		return p
	}
	return nil
}

type List struct {
	root Item
	size int
}

func (l *List) Init() *List {
	l.size = 0
	l.root.next = nil
	return l
}

func NewList() *List {
	return new(List).Init()
}

func (l *List) First() *Item {
	return l.root.next
}

func (l *List) Last() *Item {
	if l.root.next != nil {
		var p *Item
		for p = l.root.next; p.next != nil; p = p.next {
		}
		return p
	}
	return nil
}

func (l *List) Len() int {
	return l.size
}

func (l *List) insert(e, at *Item) *Item {
	e.next = at.next
	at.next = e
	e.list = l
	l.size++
	return e
}

func (l *List) Remove(e *Item) *Item {
	var p *Item
	if l.size > 0 {
		for p = &l.root; p.next != nil && p.next != e; p = p.next {
		}
		if p.next == nil {
			return nil
		}
		e = p.next
		p.next = e.next
		e.list = nil
		l.size--
		return e
	}
	return nil
}

func (l *List) insertValue(v interface{}, at *Item) *Item {
	return l.insert(&Item{Value: v}, at)
}

func (l *List) PushBack(v interface{}) *Item {
	last := l.Last()
	if last == nil {
		l.root.next = &Item{Value: v}
		l.size++
		return l.root.next
	}
	return l.insertValue(v, last)
}
