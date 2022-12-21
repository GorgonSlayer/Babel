package util

import (
	"radiola.co.nz/babel/src/worker"
)

/**

	This system is our Priority Queue implementation. It implements Heap interface, as well as Sort Interface.
	It serves to implement the behaviour on top of the Heap to track a priority queue.

	We don't touch the Heap here, but we provide the implementation. The Overseer implements the Heap behaviour.

	We mix pointers and values here, which is not good practice. However, we need to avoid duplicating items.
	Len, Swap and Less are explicit operations on two items in the queue, they are 'atomic' in the sense of not touching the queue directly.
	Push, Pop and update are directly touching the queue. We want to avoid duplicating items, so we deal in pointers.
**/

// Item /** An item in this case is a worker. We store the worker reference, the priority (unix time) and the index here**/
type Item struct {
	Priority int64 //Priority here is the Unix time. So we push the next polling time onto this priority queue, so the lowest value is the next item polled.
	Index    int   //Our index value.
	Worker   worker.Worker
}

// PriorityQueue /** PriorityQueue implements the heap Interface and holds Items. **/
type PriorityQueue []*Item

// Len /** Queue length. **/
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less /** Comparator. This works in Ascending order. **/
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority // > would give us Descending order.
}

// Swap /** Ordering Swapper **/
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i] //Swap position (external index in the list)
	pq[i].Index = i             //Index swap (item internal index)
	pq[j].Index = j
}

// Push /** This pushes our new item onto the queue. **/
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop /** This returns the last item in our priority queue. **/
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[0]
	old[0] = nil    //Memory Leak risk unless we nil this
	item.Index = -1 //Also memory leak avoidance
	*pq = old[1:n]
	return item
}

// Peek /** Peek the first item in the Queue. We want to be able to check the time of this item continuously. **/
func (pq *PriorityQueue) Peek() *Item {
	queue := *pq
	first := queue[0]
	return first
}
