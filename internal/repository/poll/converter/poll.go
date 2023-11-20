package converter

import (
	"github.com/alitvinenko/fcsempark_bot/internal/model"
	repoModel "github.com/alitvinenko/fcsempark_bot/internal/repository/poll/model"
)

func ToPollFromRepo(poll *repoModel.Poll) *model.Poll {
	return &model.Poll{
		ID:                poll.ID,
		TelegramMessageID: poll.TelegramMessageID,
		Date:              poll.Date,
		Status:            model.PollStatus(poll.Status),
		MaxPlayers:        poll.MaxPlayers,
	}
}

func ToRepoFromService(poll *model.Poll) *repoModel.Poll {
	return &repoModel.Poll{
		ID:                poll.ID,
		TelegramMessageID: poll.TelegramMessageID,
		Date:              poll.Date,
		Status:            repoModel.PollStatus(poll.Status),
		MaxPlayers:        poll.MaxPlayers,
	}
}
