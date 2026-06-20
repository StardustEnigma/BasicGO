package workers

import (
	"TaskManager/queue"
	"TaskManager/services"
	"fmt"
	"time"
)



func Worker(workerId int) {
	for taskId := range queue.TaskQueue {
		fmt.Println("Worker ",workerId,"Picked Task ",taskId)

		services.ProcessTask(taskId)
		time.Sleep(5*time.Second)

		fmt.Println("Worker", workerId, "completed task", taskId)
	}
}

