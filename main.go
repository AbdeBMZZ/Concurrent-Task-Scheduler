package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Task struct {
	ID       int
	Duration time.Duration
}

type TaskScheduler struct {
	Tasks     []Task
	Completed int
	mutex     sync.Mutex
}

func (ts *TaskScheduler) AddTask(task Task) {
	ts.Tasks = append(ts.Tasks, task)
}

func (ts *TaskScheduler) executeTask(task Task) {
	fmt.Println("Executing Task:", task.ID)
	time.Sleep(task.Duration)
	ts.mutex.Lock()
	ts.Completed++
	ts.mutex.Unlock()
	fmt.Println("Completed Task:", task.ID)
}

func (ts *TaskScheduler) Start() {
	fmt.Println("Task Scheduler Started")
	for _, task := range ts.Tasks {
		go ts.executeTask(task)
	}
}

func (ts *TaskScheduler) Stop() {
	fmt.Println("Task Scheduler Stopping...")
	for {
		ts.mutex.Lock()
		completed := ts.Completed
		ts.mutex.Unlock()

		if completed == len(ts.Tasks) {
			break
		}

		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("Task Scheduler Stopped")
}


func main() {
	taskScheduler := TaskScheduler{}

	taskScheduler.AddTask(Task{ID: 1, Duration: time.Second * 2})
	taskScheduler.AddTask(Task{ID: 2, Duration: time.Second * 3})
	taskScheduler.AddTask(Task{ID: 3, Duration: time.Second * 1})

	taskScheduler.Start()

	taskScheduler.Stop()

	fmt.Println("All tasks completed")
}
