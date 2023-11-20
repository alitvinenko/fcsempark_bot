package repository

import (
	"context"
	"github.com/alitvinenko/fcsempark_bot/internal/model"
	"time"
)

type PollRepository interface {
	Save(ctx context.Context, poll *model.Poll) error
	Delete(ctx context.Context, ID string) error
	All(ctx context.Context) ([]*model.Poll, error)
	Get(ctx context.Context, ID string) (*model.Poll, error)
	GetByDate(ctx context.Context, date time.Time) (*model.Poll, error)
}
