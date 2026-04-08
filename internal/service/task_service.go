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
func (s *TaskService) CreateTaks(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	if task.Title == "" {
		log.Println("task.Title is empty")
		return nil, domain.ErrBadRequest
	}
	err := s.repo.Save(ctx, task)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return task, nil
}
