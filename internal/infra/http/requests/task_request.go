package requests

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type FindTaskByIdRequest struct {
	Id uint64 `json:"id" validate:"required"`
}

type FindTaskByUserRequest struct {
	UserId uint64 `json:"id" validate:"required"`
}

type CreateTaskRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description" validate:"required"`
	Deadline    int64   `json:"deadline" validate:"required"`
}

type UpdateTaskRequest struct {
	Id          uint64  `json:"id" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description" validate:"required"`
	Deadline    int64   `json:"deadline" validate:"required"`
	Status      string  `json:"status" validate:"required"`
}

type DeleteTaskRequest struct {
	Id uint64 `json:"id" validate:"required"`
}

func (r CreateTaskRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Title:       r.Title,
		Description: r.Description,
		Deadline:    time.Unix(r.Deadline, 0),
	}, nil
}

func (r UpdateTaskRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Id:          r.Id,
		Title:       r.Title,
		Description: r.Description,
		Status:      domain.TaskStatus(r.Status),
		Deadline:    time.Unix(r.Deadline, 0),
	}, nil
}

func (r DeleteTaskRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Id: r.Id,
	}, nil
}

func (r FindTaskByIdRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Id: r.Id,
	}, nil
}

func (r FindTaskByUserRequest) ToDomainModel() (interface{}, error) {
	return domain.User{
		Id: r.UserId,
	}, nil
}
