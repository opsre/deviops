package scheduler

import (
	"fmt"
	"log"
	"sync"

	"dodevops-api/api/task/model"
	"dodevops-api/api/task/service"
	"dodevops-api/common"

	"github.com/robfig/cron/v3"
)

type TaskScheduler struct {
	cron         *cron.Cron
	taskService  service.ITaskAnsibleService
	scheduleJobs map[uint]cron.EntryID
	mutex        sync.RWMutex
}

func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		cron:         cron.New(cron.WithSeconds()),
		taskService:  service.NewTaskAnsibleService(common.GetDB()),
		scheduleJobs: make(map[uint]cron.EntryID),
	}
}

// Start 启动调度器
func (s *TaskScheduler) Start() error {
	log.Println("[Scheduler] Starting task scheduler service...")

	// 注册回调函数，当Service层有任务变更时通知调度器
	service.OnTaskConfigChange = func(task *model.TaskAnsible) {
		log.Printf("[Scheduler] Received task update event for Task ID: %d", task.ID)
		if err := s.AddTaskSchedule(task); err != nil {
			log.Printf("[Scheduler] Failed to update schedule for task %d: %v", task.ID, err)
		}
	}

	s.cron.Start()
	log.Println("[Scheduler] Cron engine started.")
	return s.LoadActiveSchedules()
}

// Stop 停止调度器
func (s *TaskScheduler) Stop() {
	s.cron.Stop()
}

// LoadActiveSchedules 加载所有启用的定时任务
func (s *TaskScheduler) LoadActiveSchedules() error {
	var tasks []model.TaskAnsible
	// Find recurring tasks
	if err := common.GetDB().Where("is_recurring = ?", 1).Find(&tasks).Error; err != nil {
		return err
	}

	for _, task := range tasks {
		if err := s.AddTaskSchedule(&task); err != nil {
			log.Printf("Failed to add task schedule %d: %v", task.ID, err)
		}
	}
	return nil
}

// AddTaskSchedule 添加/更新定时任务
func (s *TaskScheduler) AddTaskSchedule(task *model.TaskAnsible) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// 先移除旧的
	if entryID, exists := s.scheduleJobs[task.ID]; exists {
		s.cron.Remove(entryID)
		delete(s.scheduleJobs, task.ID)
	}

	// 检查是否应该添加
	if task.IsRecurring != 1 || task.CronExpr == "" {
		return nil
	}

	entryID, err := s.cron.AddFunc(task.CronExpr, func() {
		// 执行任务
		log.Printf("Executing scheduled task: %s (%d)", task.Name, task.ID)
		if err := s.taskService.ExecuteTask(task.ID); err != nil {
			log.Printf("Failed to execute scheduled task %d: %v", task.ID, err)
		}
	})
	if err != nil {
		return fmt.Errorf("bad cron expr '%s': %v", task.CronExpr, err)
	}

	s.scheduleJobs[task.ID] = entryID
	log.Printf("Added schedule for task %d [%s]: %s", task.ID, task.Name, task.CronExpr)
	return nil
}

// RemoveTaskSchedule 移除任务调度
func (s *TaskScheduler) RemoveTaskSchedule(taskID uint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if entryID, exists := s.scheduleJobs[taskID]; exists {
		s.cron.Remove(entryID)
		delete(s.scheduleJobs, taskID)
		log.Printf("Removed schedule for task %d", taskID)
	}
}
