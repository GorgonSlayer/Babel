package test

import (
	"github.com/stretchr/testify/assert"
	"radiola.co.nz/babel/src/overseer"
	"radiola.co.nz/babel/src/util"
	"radiola.co.nz/babel/src/util/logger"
	"radiola.co.nz/babel/src/util/mock"
	"testing"
)

/** Generates a new mock Item to insert into the Priority Queue.**/
func GenerateMockItem(priority int64, index int) *util.Item {
	item := &util.Item{
		Priority: priority,
		Index:    index,
		Worker:   mock.NewMockWorker(),
	}
	return item
}

/** This should purely test whether we generate an appropriate overseer. **/
func TestNewOverseerQueue(t *testing.T) {
	l := logger.NewLogger(false, "overseer_test.log")
	pq := overseer.NewOverseer(l)
	assert.Equal(t, 0, pq.Length())
}

/** This should test whether we push appropriately onto the heap. **/
func TestOverseerQueuePush(t *testing.T) {
	l := logger.NewLogger(false, "overseer_test.log")
	pq := overseer.NewOverseer(l)
	assert.Equal(t, 0, pq.Length())
	pq.Push(GenerateMockItem(1, 1))
}

func TestOverseerQueuePop(t *testing.T) {
	l := logger.NewLogger(false, "overseer_test.log")
	pq := overseer.NewOverseer(l)
	pq.Push(GenerateMockItem(1, 1))
	pq.Push(GenerateMockItem(2, 2))
	assert.Equal(t, 2, pq.Length())
	pq.Pop()
	assert.Equal(t, 1, pq.Length())
}

func TestOverseerQueuePeek(t *testing.T) {
	l := logger.NewLogger(false, "overseer_test.log")
	pq := overseer.NewOverseer(l)
	assert.Equal(t, 0, pq.Length())
	pq.Push(GenerateMockItem(1, 1))
	pq.Push(GenerateMockItem(2, 2))
	pq.Peek()
	assert.Equal(t, 2, pq.Length())
	assert.Equal(t, 0, pq.Peek().Index)
}

func TestOverseerQueueRemove(t *testing.T) {
	l := logger.NewLogger(false, "overseer_test.log")
	pq := overseer.NewOverseer(l)
	assert.Equal(t, 0, pq.Length())
	pq.Push(GenerateMockItem(1, 1))
	m := GenerateMockItem(2, 2)
	pq.Push(m)
	pq.Push(GenerateMockItem(3, 3))
	assert.Equal(t, 3, pq.Length())
	pq.Remove(m)
	assert.Equal(t, 2, pq.Length())
}
