package showdb

import (
	"encoding/json"
	"github.com/alitvinenko/fcsempark_bot/internal/config"
	"github.com/alitvinenko/fcsempark_bot/internal/database"
	"github.com/alitvinenko/fcsempark_bot/internal/polls/repository"
	"log"
)

func Run() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("Error on init config: %v", err)
	}

	db := database.InitDB(config.GetDatabasePath())
	defer func() { _ = db.Close() }()

	r := repository.NewPollRepository(db)

	items, err := r.All()
	if err != nil {
		log.Fatal(err)
	}

	// Выводим результаты в консоль
	log.Println("All saved items:")
	for _, p := range items {
		str, err := json.Marshal(p)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(string(str))
	}
}
