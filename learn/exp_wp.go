package main

import (
	"fmt"
	//	"time"
)

const (
	MaxJob    = 100000000
	MaxWorker = 20
)

type Job struct {
	data int
}

var (
	JobQueue = make(chan Job, MaxJob)
)

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
	Id         int
}

func (w *Worker) Start() {
	go func() {
		for {
			//first register itself
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				fmt.Println("job ", job.data, " performend by worker ", w.Id)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.quit <- true
}

func NewWorker(pool chan chan Job, id int) *Worker {
	return &Worker{
		WorkerPool: pool,
		JobChannel: make(chan Job, MaxJob),
		quit:       make(chan bool),
		Id:         id,
	}
}

type Dispatcher struct {
	pool chan chan Job
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{pool: make(chan chan Job, MaxWorker)}
}

func (d *Dispatcher) Run() {
	//start all workers
	for i := 0; i < MaxWorker; i++ {
		worker := NewWorker(d.pool, i+1)
		worker.Start()
	}

	//dispatch all jobs
	go d.Dispatch()
}

func (d *Dispatcher) Dispatch() {
	for {
		job := <-JobQueue
		go func(job Job) {
			//acquire a idle worker
			jobChannel := <-d.pool

			//Dispatch the job
			jobChannel <- job
		}(job)
	}
}

func main() {
	d := NewDispatcher()
	d.Run()

	//create job
	for i := 0; i < MaxJob; i++ {
		JobQueue <- Job{i + 1}
	}

	//	time.Sleep(time.Second * 10)
}
