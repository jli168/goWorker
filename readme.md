A Demo of Go Worker Queues Workflow
===================================

##### Description
*	The code is copied and slightly updated from Nick Saika's [Writing worker queues, in Go](http://nesv.github.io/golang/2014/02/25/worker-queues-in-go.html). It is a terrific example of showing how Go channel makes concurreny so easy.

*	It does the following:
	*	Create a `WorkRequest` struct. `WorkRequest` objects are created in `collector.go` by receiving a HTTP request, and pushed into a `WorkQueue`, which is a buffered channel.

	*	Create a `Worker` struct, which has fields of `WorkerQueue` channel and `Work` channel. `WorkerQueue` channel field is shared by all `Worker` objects, so that we can dispatch available workers through it. Each worker has a dedicated `Work` channel is used to pipeline with work request from `WorkQueue`, so that we can send work request in `WorkQueue` to this worker's `Work` channel.

	*	`dispatcher.go` creates the `WorkerQueue` channel, and create n `Worker` objects, each object will do this: whenever `WorkerQueue` channel has a space, it will put itself into the channel, and make itself able to receive work request signal and do the actual work. `dispatcher.go` also creates a goroutine to receive work requests from `WorkQueue`, and assign it to an available worker from `WorkerQueue`.

*	A little update:
	*	Originally `WorkerQueue` channel was a channel of `WorkRequest` channel, which made it a little hard to understand semantically.
	*	So I changed `WorkerQueue` channel to a channel of worker references: `var WorkerQueue chan *Worker`.
