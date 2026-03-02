package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, result chan <- int, wg *sync.WaitGroup){
	defer wg.Done()

	for task := range tasks{
		fmt.Printf("Task %d is picked up by the worker %d\n", task, id)
		time.Sleep(1 *time.Second)//simulating some logic processing
		result <- task * task
	}
}

func main(){

	const workers = 3	//no of workers
	const jobs = 7		//no of jobs

	tasks := make(chan int, jobs)		//channel to send tasks
	results := make(chan int, jobs)	//channel to receive results

	var wg sync.WaitGroup

	//spawn workers
	for i:=1;i<=workers;i++{
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	//create tasks
	for i:=1;i<=jobs;i++{
		tasks <- i
	}
	close(tasks)

	wg.Wait()
	close(results)

	//get the results
	for r :=range results{
		fmt.Println(r)
	}
}