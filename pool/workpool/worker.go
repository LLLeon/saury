package workpool

// Worker represents a
type Worker struct {
	jobChan    chan Job
	workerPool chan chan Job
}

// NewWorker returns a Worker with
func NewWorker(pool chan chan Job) *Worker {
	return &Worker{
		jobChan:    make(chan Job, 1),
		workerPool: pool,
	}
}

// Start register jobChan to the workerPool of dispatcher
// and constantly fetch job from jobChan for execution.
func (w *Worker) Start() {
	go func() {
		for {
			// register this worker's jobChan to the workerPool
			w.workerPool <- w.jobChan

			select {
			case job := <-w.jobChan:
				job.Do()
			}
		}
	}()
}
