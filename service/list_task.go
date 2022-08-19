package service

import (
	"context"
	"fmt"

	"github.com/eng-nakamura-tetsu/go-rest-api/auth"
	"github.com/eng-nakamura-tetsu/go-rest-api/entity"
	"github.com/eng-nakamura-tetsu/go-rest-api/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	ts, err := l.Repo.ListTasks(ctx, l.DB, id)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}
