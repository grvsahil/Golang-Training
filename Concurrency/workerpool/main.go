package main

import (
	"fmt"
	"sync"
	"time"
	// "time"
)

type Job struct{
	id int
	num int
}

type result struct{
	job Job
	value int
}

var jobCh = make(chan Job,10)
var resCh = make(chan result,10)

func sumOfDigits(num int) int {
	sum:=0
	for num!=0{
		digit := num % 10
		sum += digit
		num /= 10
	}
	return sum
}

func worker(wg *sync.WaitGroup)  {
	for val:=range jobCh{
		output := result{val,sumOfDigits(val.num)}
		resCh<-output
	}
	wg.Done()
}

func workerManager(noOfWorker int){
	var wg sync.WaitGroup
	for noOfWorker!=0{
		wg.Add(1)
		go worker(&wg)
		noOfWorker--;
	}
	wg.Wait()
	close(resCh)
}

func allocate(noOfJobs int){
	//Creates new jobs and send it over jobCh Channel
	for i:=1;i<=noOfJobs;i++{
		job := Job{i,2*i}
		jobCh <- job
	}
	close(jobCh)
}

func printResult(done chan bool){
	for val:=range resCh{
		fmt.Printf("Job %d: For %d , the sum of digits is %d\n\n",val.job.id,val.job.num,val.value)
	}
	done <- true
}

func main()  {
	//Record the start time
	startTime:= time.Now()

	//Lets make 100 jobs and send them over jobCh channel
	go allocate(100000000)

	done := make(chan bool)
	go printResult(done)

	//Lets make 10 workers using workerManager
	workerManager(10)

	<-done

	//Record end time
	endTime:=time.Now()
	diff:=endTime.Sub(startTime)
	fmt.Printf("Total time taken %v seconds\n",diff.Seconds())

	
}