package main

import (
	"container/list"
)

type parseTree struct {
	ts treeStack
}

type treeNode struct {
	parent   string //LHS sym
	termsym  string
	children *treeNodeQueue
}

type treeNodeQueue struct {
	queue *list.List
}

type treeStack struct {
	stack *list.List
}

//Support functions for treeNode
func (tn treeNode) String() string {
	if tn.children == nil && tn.parent == "" {
		return tn.termsym
	} else {
		return "[ " + tn.parent + " " + tn.children.String() + " ]"
	}
}

//Support functions for treeNodeQueue
func newTreeNodeQueue() treeNodeQueue {
	e := treeNodeQueue{}
	e.queue = list.New()
	return e
}

func (tq treeNodeQueue) enqueue(itm treeNode) {
	tq.queue.PushFront(itm)
}

func (tq treeNodeQueue) String() string {
	str := ""
	for e := tq.queue.Front(); e != nil; e = e.Next() {
		str += e.Value.(treeNode).String()
	}

	return str
}

//Support functions for treeStack
func newTreeStack() treeStack {
	e := treeStack{}
	e.stack = list.New()
	return e
}

func (ts treeStack) top() treeNode {
	e := ts.stack.Front()
	return e.Value.(treeNode)
}

func (ts treeStack) pop() treeNode {
	e := ts.stack.Front()
	if e != nil {
		ts.stack.Remove(e)
		return e.Value.(treeNode)
	}
	return treeNode{"", "", nil}
}

func (ts treeStack) push(node treeNode) {
	ts.stack.PushFront(node)
}

func (ts treeStack) String() string {
	str := ""
	if ts.stack.Len() > 0 {
		e := ts.stack.Back()
		str = e.Value.(treeNode).String()
		for e.Prev() != nil {
			str += e.Prev().Value.(treeNode).String()
			e = e.Prev()
		}
	}

	return str
}

//Support functions for parse tree
func (pt parseTree) String() string {
	return pt.ts.String()
}

func (pt parseTree) parseTreeShift() {
	pt.ts.push(treeNode{"", "id", nil})
}

func (pt parseTree) parseTreeReduce(LHS string, r int, op string) {
	y := newTreeNodeQueue()
	if r == 1 {
		e := pt.ts.pop()

		y.enqueue(e)
		pt.ts.push(treeNode{LHS, "", &y})
	} else {
		a := pt.ts.pop()
		b := pt.ts.pop()
		y.enqueue(a)
		y.enqueue(treeNode{"", op, nil})
		y.enqueue(b)
		pt.ts.push(treeNode{LHS, "", &y})
	}
}
