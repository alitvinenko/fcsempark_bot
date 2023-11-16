package database

import (
	"github.com/alitvinenko/fcsempark_bot/internal/polls/repository"
	"github.com/boltdb/bolt"
	"log"
)

func InitDB(path string) *bolt.DB {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.PollsBucket))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
