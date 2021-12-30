package server

import "sync/atomic"

// RingQueue 环形队列来接收消息
// 现在是无锁的，如果消息溢出，则说明达到承载上限，新的覆盖旧的。
// TODO, 待扩展特性，自动扩容，自动判断当前负载来缩容，这两个都需要锁来实现。
type RingQueue struct {
	write	int64
	read	int64
	size	int64
	queue	[]*MetricData
	noEmpty chan struct{}
}

func newRingQueue(size int32) *RingQueue {
	if size < 1 {
		panic("size must be > 1.")
	}
	rq := &RingQueue{
		write: -1,
		read: -1,
		size: int64(size),
		queue: make([]*MetricData, size),
		noEmpty: make(chan struct{}, 1),
	}
	return rq
}

func (r *RingQueue) offer(data *MetricData) {
	newWrite := atomic.AddInt64(&r.write, 1)
	index := int32(newWrite % r.size)
	r.queue[index] = data
	r.noEmpty <- struct{}{}
}

func (r *RingQueue) poll() *MetricData {
	check:
	if r.read >= r.write {
		select {
		case <-r.noEmpty:
			goto check
		}
	}
	newRead := atomic.AddInt64(&r.read, 1)
	index := int32(newRead % r.size)
	return r.queue[index]
}
