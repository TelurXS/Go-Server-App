package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type TaskService interface {
	FindById(id uint64) (domain.Task, error)
	FindAllByUserId(id uint64) ([]domain.Task, error)
	Save(t domain.Task) (domain.Task, error)
	Update(user domain.Task) (domain.Task, error)
	Delete(id uint64) error
}

type taskService struct {
	taskRepo database.TaskRepository
}

func NewTaskService(tr database.TaskRepository) TaskService {
	return taskService{
		taskRepo: tr,
	}
}

func (s taskService) FindById(id uint64) (domain.Task, error) {
	task, err := s.taskRepo.FindById(id)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return domain.Task{}, err
	}
	return task, err
}

func (s taskService) FindAllByUserId(id uint64) ([]domain.Task, error) {
	tasks, err := s.taskRepo.FindAllByUserId(id)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return []domain.Task{}, err
	}
	return tasks, err
}

func (s taskService) Save(t domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Save(t)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

func (s taskService) Update(task domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Update(task)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return domain.Task{}, err
	}

	return task, nil
}

func (s taskService) Delete(id uint64) error {
	err := s.taskRepo.Delete(id)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return err
	}

	return nil
}
