package common

var jobQueue chan func()

func EnqueueJob(job func()) {
	jobQueue <- job
}

func CreateWorker() {
	jobQueue = make(chan func(), 100)
	go func(jobs <-chan func()) {
		for job := range jobs {
			job()
		}
	}(jobQueue)
}
