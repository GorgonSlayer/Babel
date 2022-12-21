package test

import (
	"github.com/stretchr/testify/assert"
	"radiola.co.nz/babel/src/util"
	"radiola.co.nz/babel/src/util/mock"
	"testing"
)

/** This Testing should validate the Priority Queue implementation in use here. **/

var n = 10

// GeneratePriorityQueue /** Our setup and teardown function **/
func GeneratePriorityQueue() util.PriorityQueue {
	//generate a Queue of some sort, insert some mocked items.

	pq := make(util.PriorityQueue, n)
	for k := 0; k < n; k++ {
		pq[k] = &util.Item{
			Priority: int64(k),
			Index:    k,
			Worker:   mock.NewMockWorker(),
		}
	}
	return pq
}

// TestPriorityQueueLen /** Priority Queue Length validate **/
func TestPriorityQueueLen(t *testing.T) {
	pq := GeneratePriorityQueue()
	assert.Equal(t, pq.Len(), n)
}

// TestPriorityQueueLess /** Priority Queue Less validate **/
func TestPriorityQueueLess(t *testing.T) {
	pq := GeneratePriorityQueue()
	lesser := pq.Less(int(pq[1].Priority), int(pq[2].Priority))
	notLesser := pq.Less(int(pq[2].Priority), int(pq[1].Priority))
	assert.Equal(t, true, lesser)
	assert.Equal(t, false, notLesser)
}

// TestPriorityQueueSwap /** Priority Queue Swap item validate **/
func TestPriorityQueueSwap(t *testing.T) {
	pq := GeneratePriorityQueue()
	pq.Swap(1, 2)
	assert.Equal(t, pq[2].Priority, int64(1))
	assert.Equal(t, pq[1].Priority, int64(2))
}

/** We test whether or not our priority queue array accepts items appropriately. **/
func TestPriorityQueuePush(t *testing.T) {
	pq := GeneratePriorityQueue()
	i := util.Item{
		Priority: int64(n),
		Index:    n,
		Worker:   mock.NewMockWorker(),
	}
	pq.Push(&i)
	assert.Equal(t, pq[n], &i)
}

/** Verify that we do remove one item from the queue. **/
func TestPriorityQueuePop(t *testing.T) {
	pq := GeneratePriorityQueue()
	assert.Equal(t, n, pq.Len())
	j := pq[0]
	i := pq.Pop().(*util.Item)
	assert.Equal(t, n-1, pq.Len())
	assert.Equal(t, j, i) //Check that we pop the last one off.
}

/** Verify that we can view the top item. **/
func TestPriorityQueuePeek(t *testing.T) {
	pq := GeneratePriorityQueue()
	i := pq.Peek()
	j := pq[0]
	assert.Equal(t, n, pq.Len())
	assert.Equal(t, j, i) //Check that we pop the last one off.
}
