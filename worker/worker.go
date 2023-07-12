package worker

import "sync"

type workerImpl struct {
	read         chan Task
	write        chan Task
	stop         chan struct{}
	workersCount int
	countOfJobs  sync.WaitGroup
}

type Task interface {
	Execute() Task
}

func New(numberOfWorkers int, read chan Task, write chan Task, stop chan struct{}) workerImpl {
	w := workerImpl{
		read:         read,
		write:        write,
		stop:         stop,
		workersCount: numberOfWorkers,
	}
	return w
}

func (w *workerImpl) Start() {
	for i := 1; i <= w.workersCount; i++ {
		w.countOfJobs.Add(1)
		go func() {
			select {
			case task, _ := <-w.read:
				result := task.Execute()
				w.write <- result
			case <-w.stop:
				w.countOfJobs.Done()
			}
		}()
	}
	w.countOfJobs.Wait()
}
