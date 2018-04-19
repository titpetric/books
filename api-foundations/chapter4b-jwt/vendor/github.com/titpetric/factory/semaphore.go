package factory

import (
	"sync/atomic"
)

type Semaphore struct {
	semaphore int32
}

func (l *Semaphore) CanRun() bool {
	return atomic.CompareAndSwapInt32(&l.semaphore, 0, 1)
}
func (l *Semaphore) Done() {
	atomic.CompareAndSwapInt32(&l.semaphore, 1, 0)
}
