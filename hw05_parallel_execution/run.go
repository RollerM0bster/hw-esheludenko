package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var once sync.Once
	errCount := 0
	tchan := make(chan Task)
	done := make(chan struct{})

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case task, ok := <-tchan:
					if !ok {
						return
					}
					if err := task(); err != nil {
						mu.Lock()
						errCount++
						if errCount >= m {
							once.Do(func() { close(done) })
						}
						mu.Unlock()
					}
				case <-done:
					return
				}
			}
		}()
	}

	for _, task := range tasks {
		select {
		case tchan <- task:
		case <-done:
			break
		}
	}
	close(tchan)

	wg.Wait()

	mu.Lock()
	defer mu.Unlock()
	if errCount >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
