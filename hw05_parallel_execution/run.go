package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type Counter struct {
	sync.RWMutex
	val int
}

func (counter *Counter) Increment() {
	counter.Lock()
	counter.val++
	counter.Unlock()
}

func (counter *Counter) GetVal() int {
	counter.RLock()
	val := counter.val
	counter.RUnlock()
	return val
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	defer wg.Wait()

	workers := make(chan struct{}, n)
	defer close(workers)

	counter := Counter{}

	for _, task := range tasks {
		if counter.GetVal() >= m {
			return ErrErrorsLimitExceeded
		}
		wg.Add(1)
		workers <- struct{}{}
		go func(task Task, counter *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			err := task()
			if err != nil {
				counter.Increment()
			}
			<-workers
		}(task, &counter, &wg)
	}
	return nil
}
