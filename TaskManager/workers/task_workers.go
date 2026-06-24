package workers

import (
	"TaskManager/queue"
	"TaskManager/services"
	"context"
	"fmt"
	"time"
)

func Worker(workerId int) {
	for taskId := range queue.TaskQueue {
		fmt.Println("Worker ",workerId,"Picked Task ",taskId)
		ctx := context.Background()
		services.ProcessTask(taskId,ctx)
		time.Sleep(5*time.Second)

		fmt.Println("Worker", workerId, "completed task", taskId)
	}
}

