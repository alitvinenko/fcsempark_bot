package poll

import (
	"context"
	"github.com/alitvinenko/fcsempark_bot/internal/model"
)

func (s *service) AllSavedItems(ctx context.Context) ([]*model.Poll, error) {
	return s.pollRepository.All(ctx)
}
