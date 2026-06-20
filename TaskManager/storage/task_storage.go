package storage

import (
	"TaskManager/models"
	"sync"
)

var Mu sync.Mutex
var TaskData=[]models.Task{}