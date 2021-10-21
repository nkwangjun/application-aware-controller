package controller

import (
	"sync"
)

type TaskManager struct {
	sync.Mutex
	taskQueue map[string]Task
}

func (tm *TaskManager) createOrUpdate(t Task) error {
	tm.Lock()
	defer tm.Unlock()

	// TODO(wangjun): 启动AHPA任务

	return nil
}

func (tm *TaskManager) delete(id string) error {
	tm.Lock()
	defer tm.Unlock()

	// TODO(wangjun): 删除AHPA任务

	return nil
}

func NewTaskManager() *TaskManager {
	tm := &TaskManager{
		taskQueue:      make(map[string]Task),
	}

	return tm
}
