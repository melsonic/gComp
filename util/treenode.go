package util

import (
	"container/heap"
	"fmt"
)

type TreeNode struct {
	characters string
	count      int
	index      int
	left       *TreeNode
	right      *TreeNode
}

/// check if the current node is child node
func (f *TreeNode) IsEmpty() bool {
	return ((f.left == nil) && (f.right == nil))
}

/// get the encoded bit strings
/// argument: charCodes - map<huffman encoded bit sequence, character> 
/// currPath: current path from root to current node
func (f *TreeNode) EncodeNode(charCodes map[string]string, currPath string) {
	if f == nil {
		return
	}
	if f.IsEmpty() {
		charCodes[currPath] = f.characters
	}
	f.left.EncodeNode(charCodes, currPath+"0")
	f.right.EncodeNode(charCodes, currPath+"1")
}

/// Print the whole Tree rooted @TreeNode f using bfs
func (f *TreeNode) PrintTree() {
	if f == nil {
		return
	}
	var queue []*TreeNode
	queue = append(queue, f)
	for len(queue) > 0 {
		var currNode *TreeNode = queue[0]
		fmt.Println(currNode.characters)
		queue = queue[1:]
		if currNode.left != nil {
			queue = append(queue, currNode.left)
		}
		if currNode.right != nil {
			queue = append(queue, currNode.right)
		}
	}
}

type PriorityQueue []*TreeNode

/// fetch the length of the priority queue
func (pq PriorityQueue) Len() int { return len(pq) }

/// comparator for priority queue heap
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count < pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

/// push element to priority queue
/// argument: x - element to push
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*TreeNode)
	item.index = n
	*pq = append(*pq, item)
}

/// pop element from priority queue
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

/// build huffman tree from the available frequency count table 
/// argument: freqCount: character frequency count table 
/// argument: charCodes - map<huffman encoded bit sequence, character> 
func BuildTree(freqCount map[byte]int, charCodes map[string]string) *TreeNode {
	pq := make(PriorityQueue, len(freqCount))
	i := 0
	for k, v := range freqCount {
		pq[i] = &TreeNode{
			characters: string(k),
			count:      v,
			index:      i,
			left:       nil,
			right:      nil,
		}
		i++
	}
	heap.Init(&pq)

	for pq.Len() > 1 {
		node1 := heap.Pop(&pq).(*TreeNode)
		node2 := heap.Pop(&pq).(*TreeNode)
		var bindedCharacters string = node1.characters + node2.characters
		var bindedCount int = node1.count + node2.count
		newNode := TreeNode{
			characters: bindedCharacters,
			count:      bindedCount,
			left:       node1,
			right:      node2,
		}
		heap.Push(&pq, &newNode)
	}
	rootNode := heap.Pop(&pq).(*TreeNode)
	return rootNode
}
