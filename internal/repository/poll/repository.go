package poll

import (
	"context"
	"errors"
	"github.com/alitvinenko/fcsempark_bot/internal/model"
	def "github.com/alitvinenko/fcsempark_bot/internal/repository"
	"github.com/alitvinenko/fcsempark_bot/internal/repository/poll/converter"
	repoModel "github.com/alitvinenko/fcsempark_bot/internal/repository/poll/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
	"time"
)

var _ def.PollRepository = (*repository)(nil)

type repository struct {
	db *gorm.DB
	m  sync.RWMutex
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(ctx context.Context, poll *model.Poll) error {
	r.m.Lock()
	defer r.m.Unlock()

	repoPoll := converter.ToRepoFromService(poll)
	result := r.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&repoPoll)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, ID string) error {
	r.m.Lock()
	defer r.m.Unlock()

	var poll repoModel.Poll
	r.db.Delete(&poll, ID)

	return nil
}

func (r *repository) All(ctx context.Context) ([]*model.Poll, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var repoPolls []*repoModel.Poll

	result := r.db.Find(&repoPolls)
	if result.Error != nil {
		return nil, result.Error
	}

	var polls []*model.Poll
	for _, poll := range repoPolls {
		polls = append(polls, converter.ToPollFromRepo(poll))
	}

	return polls, nil
}

func (r *repository) Get(ctx context.Context, ID string) (*model.Poll, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var poll repoModel.Poll
	result := r.db.First(&poll, "id = ?", ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return converter.ToPollFromRepo(&poll), nil
}

func (r *repository) GetByDate(ctx context.Context, date time.Time) (*model.Poll, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var poll repoModel.Poll
	result := r.db.First(&poll, "date = ?", date)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return converter.ToPollFromRepo(&poll), nil
}
