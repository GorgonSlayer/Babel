package overseer

import (
	"container/heap"
	"go.uber.org/zap"
	"net/http"
	"radiola.co.nz/babel/src/util"
	"radiola.co.nz/babel/src/util/logger"
	"radiola.co.nz/babel/src/worker"
	"time"
)

/**
	Overseer is the task queue system. It 'Oversees' the tasks which workers have.
**/

// IOverseer /** Overseer interface for what behaviour the Overseer should implement. **/
type IOverseer interface {
	Push(item *util.Item)
	Pop() *util.Item
	Peek() *util.Item
	Update(item *util.Item, worker worker.Worker, priority int64)
	Remove(item *util.Item)
	OverseerLoop()
	workerProcess(*util.Item)
}

// Overseer /** Overseer Struct **/
type Overseer struct {
	queue  *util.PriorityQueue
	logger logger.Logger
}

func NewOverseer(logger logger.Logger) Overseer {
	return Overseer{
		queue:  NewPriorityQueue(),
		logger: logger,
	}
}

// NewPriorityQueue /** New Priority Queue **/
func NewPriorityQueue() *util.PriorityQueue {
	pq := make(util.PriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

// Length /** Plain wrapper around the Array. Primarily used for testing. **/
func (o Overseer) Length() int {
	return o.queue.Len()
}

// Push /** Pushes an item onto the Heap. **/
func (o Overseer) Push(item *util.Item) {
	item.Priority = item.Priority + item.Worker.RefreshRate()
	heap.Push(o.queue, item)
	o.logger.Zap.Info("Pushed item onto Priority Queue", zap.Any("item", item))
}

// Pop /** Pops an item off the Heap. **/
func (o Overseer) Pop() *util.Item {
	item := heap.Pop(o.queue).(*util.Item)
	o.logger.Zap.Info("Popped Item off Priority Queue", zap.Any("item", item))
	return item
}

// Peek /** Peek the top item on the Heap. **/
func (o Overseer) Peek() *util.Item {
	item := o.queue.Peek()
	o.logger.Zap.Info("Peeked Item on Priority Queue", zap.Any("item", item))
	return item
}

// Update /** Update is an internal method to modify the priority and value of items in the queue.**/
func (o *Overseer) Update(item *util.Item, worker worker.Worker, priority int64) {
	item.Worker = worker
	item.Priority = priority
	heap.Fix(o.queue, item.Index)
}

// Remove /** Removes the item from the priority queue. **/
func (o *Overseer) Remove(item *util.Item) {
	//Implements removing items from the queue
	heap.Remove(o.queue, item.Index)
}

// OverseerLoop /** Primary loop for this OverSeer system. **/
func (o Overseer) OverseerLoop() {
	for {
		item := o.Peek()
		if item.Priority < time.Now().Unix() {
			i := o.Pop()
			go o.workerProcess(i)
			i.Priority = i.Priority + i.Worker.RefreshRate() //Add whatever extra time on until we poll again.
			o.Push(i)
		}
	}
}

/** This is our concurrent process. We spin this off for each request we fire off. **/
func (o Overseer) workerProcess(item *util.Item) bool {
	client := http.Client{}
	res, err := item.Worker.IntakeRequest(&client)
	if err != nil {
		o.logger.Zap.Error("Worker Process inside GoRoutine incurred an Error during IntakeRequest", zap.Any("Error", err))
		return false
	}
	tce, err := item.Worker.ProcessData(res)
	if err != nil {
		o.logger.Zap.Error("Worker Process inside GoRoutine incurred an Error during ProcessData", zap.Any("Error", err))
		return false
	}
	success, err := item.Worker.OuttakeRequest(&client, tce)
	if err != nil {
		o.logger.Zap.Error("Worker Process inside GoRoutine incurred an Error during OuttakeRequest", zap.Any("Error", err))
		return false
	}
	return success
}
