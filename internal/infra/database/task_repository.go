package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const TasksTableName = "tasks"

type task struct {
	Id          uint64            `db:"id,omitempty"`
	UserId      uint64            `db:"user_id"`
	Title       string            `db:"title"`
	Description *string           `db:"description"`
	Deadline    time.Time         `db:"deadline"`
	Status      domain.TaskStatus `db:"status"`
	CreatedDate time.Time         `db:"created_date"`
	UpdatedDate time.Time         `db:"updated_date"`
	DeletedDate *time.Time        `db:"deleted_date"`
}

type TaskRepository interface {
	FindById(id uint64) (domain.Task, error)
	FindAllByUserId(id uint64) ([]domain.Task, error)
	Save(t domain.Task) (domain.Task, error)
	Update(user domain.Task) (domain.Task, error)
	Delete(id uint64) error
}

type taskRepository struct {
	coll db.Collection
	sess db.Session
}

func NewTaskRepository(session db.Session) TaskRepository {
	return taskRepository{
		coll: session.Collection(TasksTableName),
		sess: session,
	}
}

func (r taskRepository) FindAllByUserId(id uint64) ([]domain.Task, error) {
	var tasks []task
	err := r.coll.Find(db.Cond{"user_id": id, "deleted_date": nil}).All(&tasks)
	if err != nil {
		return []domain.Task{}, err
	}

	var model_tasks []domain.Task

	for _, task := range tasks {
		model_tasks = append(model_tasks, r.mapModelToDomain(task))
	}

	return model_tasks, nil
}

func (r taskRepository) FindById(id uint64) (domain.Task, error) {
	var task task
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&task)
	if err != nil {
		return domain.Task{}, err
	}

	return r.mapModelToDomain(task), nil
}

func (r taskRepository) Save(t domain.Task) (domain.Task, error) {
	tsk := r.mapDomainToModel(t)
	tsk.CreatedDate = time.Now()
	tsk.UpdatedDate = time.Now()
	err := r.coll.InsertReturning(&tsk)
	if err != nil {
		return domain.Task{}, err
	}
	t = r.mapModelToDomain(tsk)
	return t, nil
}

func (r taskRepository) Update(t domain.Task) (domain.Task, error) {
	u := r.mapDomainToModel(t)
	u.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": u.Id, "deleted_date": nil}).Update(&u)
	if err != nil {
		return domain.Task{}, err
	}
	return r.mapModelToDomain(u), nil
}

func (r taskRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r taskRepository) mapDomainToModel(t domain.Task) task {
	return task{
		Id:          t.Id,
		UserId:      t.UserId,
		Title:       t.Title,
		Description: t.Description,
		Deadline:    t.Deadline,
		Status:      t.Status,
		CreatedDate: t.CreatedDate,
		UpdatedDate: t.UpdatedDate,
		DeletedDate: t.DeletedDate,
	}
}

func (r taskRepository) mapModelToDomain(t task) domain.Task {
	return domain.Task{
		Id:          t.Id,
		UserId:      t.UserId,
		Title:       t.Title,
		Description: t.Description,
		Deadline:    t.Deadline,
		Status:      t.Status,
		CreatedDate: t.CreatedDate,
		UpdatedDate: t.UpdatedDate,
		DeletedDate: t.DeletedDate,
	}
}
