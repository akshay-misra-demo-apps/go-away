package workerpool

import (
	"fmt"
)

func worker(id int, job <-chan int, result chan<- int) {
	for i := range job {
		fmt.Println("processing job", id)
		i = i * 2
		result <- i
	}
}

func ProduceJobs() {
	numberOfWorkers := 10
	jobs := make(chan int)
	results := make(chan int)

	for i := 0; i < numberOfWorkers; i++ {
		go worker(i, jobs, results)
	}

	numberOfJobs := 100

	for i := 0; i < numberOfJobs; i++ {
		jobs <- i // produce job
	}
	close(jobs)

	for i := 0; i < numberOfJobs; i++ {
		<-results
	}
	close(results)
}
