package repository

import (
	"apiGo/internal/domain"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(ctx context.Context, task *domain.Task) error {
	query := `INSERT INTO tasks(id, title, done, created_at) VALUES( $1, $2, $3, $4)`

	_, err := r.db.Exec(ctx, query, task.ID, task.Title, task.Done, task.CreatedAt)
	if err != nil {
		return fmt.Errorf("Save task:  %w", err)
	}
	return nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (task *domain.Task, err error) {
	query := `SELECT * FROM tasks WHERE id = $1`

	task = &domain.Task{}
	err = r.db.QueryRow(ctx, query, id).Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt)

	if errors.Is(err, pgx.ErrNoRows) {

		return nil, fmt.Errorf("FindByID: %w", domain.ErrTaskNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("FindByID: %w", err)
	}
	return task, nil
}

func (r *Repository) FindAll(ctx context.Context) ([]*domain.Task, error) {
	query := `SELECT * FROM tasks`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("FindAll tasks: %w", err)
	}

	defer rows.Close()

	var tasks []*domain.Task

	for rows.Next() {
		task := &domain.Task{}

		err := rows.Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt)

		if err != nil {
			return nil, fmt.Errorf("FindAll tasks: %w", err)
		}

		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FindAll tasks: %w", err)
	}
	return tasks, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM tasks WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("Delete task: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("Delete task: %w", domain.ErrTaskNotFound)
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, task *domain.Task) error {
	query := `UPDATE tasks SET title = $1, done = $2 WHERE id = $3`
	result, err := r.db.Exec(ctx, query, task.Title, task.Done, task.ID)
	if err != nil {
		return fmt.Errorf("Update task: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("Update task: %w", domain.ErrTaskNotFound)
	}
	return nil
}
