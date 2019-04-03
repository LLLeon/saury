package workpool

import "runtime"

var (
	// numCPU represents the number of logical CPUs
	// usable by the current process.
	numCPU = runtime.NumCPU()
)

// Dispatcher is used to distribute jobs to workers.
type Dispatcher struct {
	jobQueue chan Job

	// used to pass Job between workers and dispatcher
	workerPool chan chan Job
}

// NewDispatcher returns a Dispatcher with fixed number
// of workers.
func NewDispatcher(maxWorkers int) *Dispatcher {
	if maxWorkers <= 0 || maxWorkers > numCPU {
		maxWorkers = numCPU
	}

	return &Dispatcher{
		jobQueue:   make(chan Job, 1),
		workerPool: make(chan chan Job, maxWorkers),
	}
}

// Run the Dispatcher, start a fixed number of workers
// and keep assigning them tasks.
func (d *Dispatcher) Run() {
	for i := 0; i < numCPU; i++ {
		worker := NewWorker(d.workerPool)
		worker.Start()
	}

	go d.dispatch()
}

// dispatch constantly pull job from jobQueue and put it
// into jobChan for the corresponding worker to execute.
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			jobChan := <-d.workerPool
			jobChan <- job
		}
	}
}

// PushToJobQ put the job in the job queue.
func (d *Dispatcher) PushToJobQ(job Job) {
	d.jobQueue <- job
}
