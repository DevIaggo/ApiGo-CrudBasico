package domain

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, task *Task) error
	FindByID(ctx context.Context, id string) (*Task, error)
	FindAll(ctx context.Context) ([]*Task, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, task *Task, id string) error
}
