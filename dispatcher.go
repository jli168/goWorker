package main

import "fmt"

var WorkerQueue chan chan WorkRequest

func StartDispatcher(nworkers int) {
  // First, initialize the channel we are going to but the workers' work channels into.
  WorkerQueue = make(chan chan WorkRequest, nworkers)
  
  // Now, create all of our workers.
  for i := 0; i<nworkers; i++ {
    fmt.Println("Starting worker", i+1)
    worker := NewWorker(i+1, WorkerQueue)
    worker.Start()
  }
  
  go func() {
    for {
      select {
      // Note: once we have `workRequest` in `WorkQueue`, we will be able to receive data from it
      case work := <-WorkQueue:
        fmt.Println("Received work requeust")
        go func() {
          // Note: notice it is a different channel name, not `WorkQueue`!
          // here we receive a worker from `WorkerQueue`, as long as there is one there, otherwise it is blocked
          worker := <-WorkerQueue
          
          fmt.Println("Dispatching work request")
          // Note: send `work` to `worker`, which is a `workRequest` channel
          // so that in `worker.go`, it can start working on `case work := <-w.Work:`
          worker <- work
        }()
      }
    }
  }()
}
