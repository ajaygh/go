package main

import (
	"fmt"

	"github.com/jrallison/go-workers"
)

func myJob(msg *workers.Msg) {
	fmt.Println("message jobid = ", msg.Jid())
}

func Add(msg *workers.Msg) {
	fmt.Println("messgage", msg.Jid())
}

type myMiddleware struct{}

func (r *myMiddleware) Call(queue string, message *workers.Msg, next func() bool) (acknowledge bool) {
	//do something before each message is processed
	fmt.Println("jid ", message.Jid())
	acknowledge = next()
	//do something after each message is processed
	return
}

func main() {
	workers.Configure(map[string]string{
		//location of database
		"server": "localhost:6379",
		//instance of the database
		"database": "0",
		//number of connections to keep open with redis
		"pool": "30",
		//unique process id
		"process": "1",
	})

	workers.Middleware.Append(&myMiddleware{})

	//Pull messages from myqueue with concurrency of 10
	workers.Process("myqueue", myJob, 10)

	//Pull messages from myqueue2 with concurrency of 20
	workers.Process("myqueue2", myJob, 20)

	//Add a job to the queue
	workers.Enqueue("myqueue3", "Add", []int{1, 2})

	//Add a job to the queue with retry
	workers.EnqueueWithOptions("myqueue3", "Add", []int{1, 2}, workers.EnqueueOptions{Retry: true})

	//See stats on localhost:8080
	//	workers.StatsServer(8085)

	workers.Run()
}
