package controllers

import (
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type TaskController struct {
	taskService app.TaskService
}

func NewTaskController(ts app.TaskService) TaskController {
	return TaskController{
		taskService: ts,
	}
}

func (c TaskController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		task, err := requests.Bind(r, requests.FindTaskByIdRequest{}, domain.Task{})

		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		task, err = c.taskService.FindById(task.Id)

		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var taskDto resources.TaskDto
		Success(w, taskDto.DomainToDto(task))
	}
}

func (c TaskController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := r.Context().Value(UserKey).(domain.User)
		task, err := requests.Bind(r, requests.CreateTaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		task.UserId = user.Id
		task.Status = domain.New
		task, err = c.taskService.Save(task)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}
		var tDto resources.TaskDto
		tDto = tDto.DomainToDto(task)
		Created(w, tDto)
	}
}

func (c TaskController) FindByAuthUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := r.Context().Value(UserKey).(domain.User)

		tasks, err := c.taskService.FindAllByUserId(user.Id)

		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var taskDto resources.TaskDto
		var dto_tasks []resources.TaskDto

		for _, task := range tasks {
			dto_tasks = append(dto_tasks, taskDto.DomainToDto(task))
		}

		Success(w, dto_tasks)
	}
}

func (c TaskController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		task, err := requests.Bind(r, requests.UpdateTaskRequest{}, domain.Task{})

		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		found, err := c.taskService.FindById(task.Id)

		if err != nil {
			log.Printf("TaskController: %s", err)
			NotFound(w, err)
			return
		}

		found.Title = task.Title
		found.Description = task.Description
		found.Deadline = task.Deadline
		found.Status = task.Status
		task, err = c.taskService.Update(found)

		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var taskDto resources.TaskDto
		Success(w, taskDto.DomainToDto(task))
	}
}

func (c TaskController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task, err := requests.Bind(r, requests.DeleteTaskRequest{}, domain.Task{})

		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		err = c.taskService.Delete(task.Id)

		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
