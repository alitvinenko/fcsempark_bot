package repository

import (
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"time"
)

var ErrStopIteration = errors.New("stop iteration")

type PollStatus int

const (
	PollStatusActive PollStatus = iota
	PollStatusClosed
)

type Poll struct {
	ID          string     `json:"poll_id"`
	TgMessageID int        `json:"tg_message_id"`
	Day         time.Time  `json:"date"`
	Status      PollStatus `json:"status"`
	MaxPlayers  int        `json:"max_players"`
}

const PollsBucket = "Polls"

type PollRepository struct {
	db *bolt.DB
}

func NewPollRepository(db *bolt.DB) *PollRepository {
	return &PollRepository{db: db}
}

func (r *PollRepository) Save(poll Poll) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(PollsBucket))

		data, err := json.Marshal(poll)
		if err != nil {
			return err
		}

		return b.Put([]byte(poll.ID), data)
	})
}

func (r *PollRepository) All() ([]Poll, error) {
	var polls []Poll

	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(PollsBucket))

		return b.ForEach(func(k, v []byte) error {
			var poll Poll

			err := json.Unmarshal(v, &poll)
			if err != nil {
				return err
			}

			polls = append(polls, poll)

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return polls, nil
}

func (r *PollRepository) IsExists(gameDay time.Time) (bool, error) {
	poll, err := r.getOneByFilter(func(p Poll) bool {
		return p.Day.Equal(gameDay)
	})
	if err != nil && !errors.Is(err, ErrStopIteration) {
		return false, err
	}

	return poll.ID != "", nil
}

func (r *PollRepository) ByID(ID string) (Poll, error) {
	poll, err := r.getOneByFilter(func(p Poll) bool {
		return p.ID == ID
	})
	if err != nil && !errors.Is(err, ErrStopIteration) {
		return Poll{}, err
	}

	return poll, nil
}

func (r *PollRepository) getOneByFilter(filter func(Poll) bool) (Poll, error) {
	var poll Poll

	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(PollsBucket))

		return b.ForEach(func(k, v []byte) error {
			// Декодируем JSON
			var p Poll
			err := json.Unmarshal(v, &p)
			if err != nil {
				return err
			}

			if filter(p) {
				poll = p
				return ErrStopIteration
			}

			return nil
		})
	})
	if err != nil && !errors.Is(err, ErrStopIteration) {
		return Poll{}, err
	}

	return poll, nil
}

func (r *PollRepository) getByFilter(filter func(Poll) bool) ([]Poll, error) {
	var polls []Poll

	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(PollsBucket))

		return b.ForEach(func(k, v []byte) error {
			var poll Poll

			err := json.Unmarshal(v, &poll)
			if err != nil {
				return err
			}

			if filter(poll) {
				polls = append(polls, poll)
			}

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return polls, nil
}

func (r *PollRepository) delete(id string) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(PollsBucket))
		return b.Delete([]byte(id))
	})
}
