package main

import "fmt"

var WorkerQueue chan *Worker

func StartDispatcher(nworkers int) {
  // First, initialize the channel we are going to put the workers into.
  WorkerQueue = make(chan *Worker, nworkers)
  
  // Now, create all of our workers.
  for i := 0; i<nworkers; i++ {
    fmt.Println("Starting worker", i+1)
    worker := NewWorker(i+1, WorkerQueue)
    worker.Start()
  }
  
  go func() {
    for {
      select {
      // Note: once we have `workRequest` in `WorkQueue` from `collector.go`,
      // this line wil be unblocked and we will be able to receive data from it.
      case work := <-WorkQueue:
        fmt.Println("Received work requeust")
        go func() {
          // Note: notice it is a different channel name `WorkerQueue`, not `WorkQueue`!
          // here we receive a worker from `WorkerQueue` channel.
          worker := <-WorkerQueue
          
          fmt.Println("Dispatching work request")
          // Note: send `work` to `worker.Work`, which is a `workRequest` channel
          // so that in `worker.go`, it will unblock `case work := <-w.Work:`
          worker.Work <- work
        }()
      }
    }
  }()
}
