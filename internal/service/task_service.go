package service

import (
	"apiGo/internal/domain"
	"context"
	"log"
)

type TaskService struct {
	repo domain.Repository
}

func NewTaskService(repo domain.Repository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, task *domain.Task) error {
	if task.Title == "" {
		log.Println("task.Title is empty")
		return domain.ErrBadRequest
	}
	if err := s.repo.Save(ctx, task); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := s.repo.FindAll(ctx)
	if err != nil {
		log.Printf("error: %s", err)
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTaskById(ctx context.Context, id string) (*domain.Task, error) {
	if id == "" {
		log.Printf("error: %s", domain.ErrBadRequest.Error())
		return nil, domain.ErrBadRequest
	}
	task, err := s.repo.FindByID(ctx, id)
	if err != nil {
		log.Printf("error: %s", err)
		return nil, err
	}
	return task, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	if id == "" {
		log.Printf("error: %s", domain.ErrBadRequest.Error())
		return domain.ErrBadRequest
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		log.Printf("error: %s", err)
		return err
	}
	return nil
}

func (s *TaskService) UpdateTask(ctx context.Context, task *domain.Task) error {
	if task.Title == "" {
		log.Printf("error: %s", domain.ErrBadRequest.Error())
		return domain.ErrBadRequest
	}
	if err := s.repo.Update(ctx, task); err != nil {
		log.Printf("error: %s", err)
		return err
	}
	return nil
}
