// double linked list implementation

package main

import "fmt"

type node struct {
	data int
	prev *node
	next *node 
}

func main() {
	var n *node
	var start *node 
	var end *node 
	// var searchKey int = 50
	// var result *node 
	// var totalNodes int 
	// n = newNode1(10)
	// start, end = addAtStart(start, end, n)
	// n = newNode1(20)
	// start, end = addAtStart(start, end, n)
	// n = newNode1(30)
	// start, end = addAtStart(start, end, n)
	// n = newNode1(50)
	// start, end = addAtStart(start, end, n)
	// n = newNode1(60)
	// start, end = addAtLast(start, end, n)
	// fmt.Println("Forward traversal: ")
	// traverseforward(start)
	// fmt.Printf("Backward traversal: ")
	// traversebackward(end)
	// result = searchNode(start, searchKey)
	// fmt.Printf("Node %d found at %v\n", searchKey, result)
	// totalNodes = countNodes(start)
	// fmt.Printf("Total number of Nodes in Linked List: %d\n", totalNodes)
	// start, end = delete(start, end, result)
	// traverseforward(start)
	// traversebackward(end)
	n = newNode1(10)
	start, end = insertSortedNode(start, end, n)
	n = newNode1(20)
	start, end = insertSortedNode(start, end, n)
	n = newNode1(50)
	start, end = insertSortedNode(start, end, n)
	n = newNode1(5)
	start, end = insertSortedNode(start, end, n)
	n = newNode1(15)
	start, end = insertSortedNode(start, end, n)
	traverseforward(start)
	traversebackward(end)
}

func newNode1(k int) *node{
	n := &node{k, nil,nil}
	return n 
}

func addAtStart(start *node, end *node, n *node) (*node, *node) {
	var p *node 
	if start != nil {
		n.next = start
		start.prev = n
	}
	start = n
	for p=start;p.next!=nil;p=p.next{
	}
	end = p
	
	return start, end 
}
// insert at end
func addAtLast(start *node, end *node, n*node) (*node, *node) {
	var p *node 
	if end != nil {
		p = end 
		p.next = n 
		n.prev = p 
		end = n 
	} else {
		start = end
	}
	return start, end 
}

func traverseforward(start *node) {
	var p *node 
	for p=start;p!=nil;p=p.next{
		displayNode(p)
	}
	fmt.Printf("\n")
}

func displayNode(p *node) {
	fmt.Printf("%d ", p.data)
}

func traversebackward(end *node) {
	var p *node 
	for p=end;p!=nil;p=p.prev {
		displayNode(p)
	}
	fmt.Printf("\n")
}

func searchNode(start *node, data int)  *node {
	var p *node 
	for p=start;p!=nil;p=p.next {
		if p.data == data {
			return p 
		}
	}
	return nil
}

func countNodes(start *node) int {
	var ctr int = 0
	for p:=start; p!=nil;p=p.next {
		ctr = ctr + 1
		
	}
	return ctr 
}


func delete(start *node, end *node, n *node) (*node, *node) {
	var t *node 
	var p *node
	if start == end {
		start = nil 
		end = nil 
	} else if end == n {
		end = end.prev 
		end.next = nil 
		n.prev = nil 
	} else if start == n {
		start  = start.next
		start.prev = nil
		n.next = nil
	}else {
		for p=start;p!=nil;p=p.next {
			t = p.prev
			if p == n  {
			t.next = n.next
			t = n.next
			t.prev = n.prev
			n.next = nil
			n.prev = nil 
			break  
		   }
		}

	}
	
	return start, end
}

func insertSortedNode(start *node, end *node, n*node) (*node, *node) {
	var p *node 
	var t *node 
	if start == nil {
		n.next = nil
		n.prev = nil
		start = n 
		end = n
	} else {
		for p=start; p!=nil;p=p.next {
			if p.data > n.data {
				if p == start {
					n.next = start
					start.prev = n
					n.prev = nil  
					start = n 
					break 
			 	} else {
					t = p.prev
					t.next = n
					n.next = p
					n.prev = t
					p.prev = n
					break
		   		}
			}

		}
		if p == nil {
			end.next = n
			n.prev = end
			n.next = nil    
			end = n
		}
	}
	return start, end 
}