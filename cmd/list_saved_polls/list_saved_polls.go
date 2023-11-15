package main

import (
	"encoding/json"
	"fmt"
	"github.com/alitvinenko/fcsempark_bot/internal/polls/repository"
	"github.com/boltdb/bolt"
	"log"
)

const dbFilePath = "data/db.json"

func main() {
	db := initDB(dbFilePath)
	defer func() { _ = db.Close() }()

	r := repository.NewPollRepository(db)

	items, err := r.All()
	if err != nil {
		log.Fatal(err)
	}

	// Выводим результаты в консоль
	fmt.Println("All saved items:")
	for _, p := range items {
		str, err := json.Marshal(p)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf(string(str) + "\n")
	}
}

func initDB(path string) *bolt.DB {
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
