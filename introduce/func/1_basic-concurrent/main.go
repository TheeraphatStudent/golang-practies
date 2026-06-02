package main

import (
	"fmt"
	"sync"
	"time"
)

func RunWorkers(numWorkers int, numJobs int) {
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	for id := 1; id <= numWorkers; id++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", workerID, job)
				time.Sleep(1 * time.Second)
			}
		}(id)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
}

func main() {
	RunWorkers(3, 9)
}
