package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	for i:=0; i < len(tasks); i ++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			tasks[i-1]();
		}()
	}
	wg.Wait()
	return nil
}
