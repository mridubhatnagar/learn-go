// single linked list implementation

package main 

import "fmt"

type Node struct {
	data int
	next *Node  
}

func main() {
	var n *Node
	var start *Node = nil 
	// var result *Node 
	// var searchKey int = 30
	// n = newNode(10)
	// start = addAtBeginning(start, n)
	// n = newNode(20)
	// start = addAtBeginning(start, n)
	// // traverse(start)
	// n = newNode(30)
	// start = addAtBeginning(start, n)
	// // traverse(start)
	// n = newNode(50)
	// start = addAtEnd(start, n)
	// traverse(start)
	// result = search(start, searchKey)
	// fmt.Printf("Key: %d found at %v\n", searchKey, result)
	// count := nodeCount(start)
	// fmt.Printf("Total nodes in linked list are: %d\n",count)
	// start = deleteNode(start, result)
	// traverse(start)
	n = newNode(20)
	start = insertSorted(start, n)
	traverse(start)
	n = newNode(10)
	start = insertSorted(start, n)
	// start = insertSorted(start, n)
	// n = newNode(25)
	// start = insertSorted(start, n)
	// n = newNode(30)
	// start = insertSorted(start, n)
	// n = newNode(5)
	// start = insertSorted(start, n)
	// n = newNode(1)
	// start = insertSorted(start, n)
	// n = newNode(55)
	// start = insertSorted(start, n)
	// n = newNode(15)
	// start = insertSorted(start, n)
	traverse(start)
	
}

func newNode(d int) *Node{
	var k *Node 
	k = &Node{d, nil}
	return k 

}

func addAtBeginning(start *Node, n *Node) *Node {
	n.next = start 
	start = n 
	return start
}

func traverse(start *Node) {
	for p:=start;p != nil;p=p.next {
		display(p)
	} 
	fmt.Printf("\n")
}

func display(n *Node) {
	fmt.Printf("%d ", n.data)
}

func addAtEnd(start *Node, n *Node) *Node{
	
	var p *Node 
	if start == nil {
		start = n 
	} else {
		for p=start;p.next != nil;p=p.next {}
		p.next = n 
	}
	return start 
}

// TODO

// 1. search(data) return Node or nil. 
// 2. delete(Node) return start.
// 3. insertSorted(Node) return start. 
// 4. nodeCount(start) return count of nodes. 

func search(start *Node, data int) *Node {
	var p *Node 
	for p=start;p!=nil;p=p.next {
		if p.data == data {
			return p 
		}

	}
	return nil
}

func nodeCount(start *Node) int{
	var ctr int = 0
	for p:=start;p!=nil;p=p.next {
		ctr = ctr + 1
	}
	return ctr 
}


func deleteNode(start *Node, n *Node) *Node{
	var p *Node
	// deletes first node
	if start == n {
		start = start.next
		n.next = nil  
		return start
	}
	for p=start; p != nil; p=p.next {
		// deletes any node from middle
        if p.next == n  {
            p.next = n.next
            n.next = nil
			break
	    }  
	}
	return start 
}

// lock step traversal
func insertSorted(start *Node, n *Node) *Node{
	var p *Node 
	var prev *Node 
	
	for p = start; p!=nil;  {
		if p.data > n.data {
			break
		}
		prev = p
		p = p.next
	}
	if p == start {
		n.next = start
		start = n
	} else {
		prev.next = n
		n.next = p
	}
	return start
}