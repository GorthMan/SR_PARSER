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
	} else if tn.children != nil {
		return "[ " + tn.parent + " " + tn.termsym + " ]"
	} else {
		return "[ " + tn.parent + " " + tn.children.queue.Front().Value.(treeNode).String() + " ]"
	}
}

//Support functions for treeStack
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
func (pt parseTree) parseTreeShift() {
	pt.ts.push(treeNode{"", "id", nil})
}

func (pt parseTree) parseTreeReduce(LHS string, r int) {
	if r == 1 {
		e := pt.ts.pop()
		x := treeNode{LHS, "", nil}
		x.children.queue.Front(e)
		pt.ts.push(treeNode{LHS, "", x})
	} else {

	}
}
