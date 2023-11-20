package service

import (
	"context"
	"github.com/alitvinenko/fcsempark_bot/internal/model"
)

type PollService interface {
	CreateAndPin(ctx context.Context) error
	CheckAndClose(ctx context.Context, pollID string, votedCount int) error
	AllSavedItems(ctx context.Context) ([]*model.Poll, error)
}
