package linkedlist

import (
	"fmt"
	"log"
	"strings"
)

type Element struct {
	parentList *LinkedList
	left       *Element
	right      *Element
	key        string
	value      string
	isRoot     bool
}

type LinkedList struct {
	list *Element
	head *Element
	tail *Element
}

func New() *LinkedList {
	e := &Element{
		isRoot: true,
	}
	ll := &LinkedList{
		list: e,
		head: e,
		tail: e,
	}
	e.parentList = ll
	return ll
}

func (e *Element) Next() *Element {
	return e.right
}

func (e *Element) Prev() *Element {
	return e.left
}

func (e *Element) InsertNext(key, value string) *Element {
	el := &Element{
		parentList: e.parentList,
		left:       e,
		right:      e.right,
		key:        key,
		value:      value,
		isRoot:     false,
	}
	e.right = el
	// update the right node to point this new left node
	if el.right != nil {
		el.right.left = el
	}
	return el
}

func (e *Element) IsRoot() bool {
	return e.isRoot
}

func (e *Element) Value() string {
	return e.value
}

func (e *Element) Key() string {
	return e.key
}

func (e *Element) Delete() {
	if e.left == nil && e.right == nil {
		log.Printf("node already deleted %#v", e)
		return
	}
	if e.left == nil {
		log.Printf("deleting root? don't do that. isRoot? %v", e.isRoot)
		return
	}
	e.left.right = e.right

	if e.right == nil {
		// don't lose track of the tail node
		e.parentList.tail = e.left
	} else {
		e.right.left = e.left
	}

	// clear out references from this node
	e.left, e.right = nil, nil
	e.value += " evicted" // debug, should be inaccessible now
}

func (l *LinkedList) Last() *Element {
	if l.tail.left == nil {
		// we deleted the last element, and the tail is lost
		// do we link back to the parent list in the element to be able to track this?
		// for now, reset tail
		l.tail = l.head
		// log.Println("reset tail to head")
	}
	for l.tail.right != nil {
		// log.Printf("%#v is not tail, checking next", l.tail)
		l.tail = l.tail.right
	}
	if l.tail == nil {
		// log.Println("wtf, how did this get to nil")
	}
	// log.Printf("found new tail: %#v", l.tail)
	return l.tail
}

func (l *LinkedList) First() *Element {
	return l.head
}

func (l *LinkedList) String() string {
	a := &Element{}
	*a = *l.First()
	var list []string
	for {
		list = append(list, fmt.Sprintf("%v[%v]", a.key, a.value))
		if a.right == nil {
			break
		}
		a = a.right
	}
	return strings.Join(list, " :: ")
}
