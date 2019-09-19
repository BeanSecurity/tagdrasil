package main

import (
	"tagdrasil/controlers"
	"tagdrasil/repository"
	"tagdrasil/manager"
	"os"
	"database/sql"
	_ "github.com/bmizerany/pq"
	"log"
)

func main() {
	//init repo
	// dburl := os.Getenv("DATABASE_URL")
	dsn := os.Getenv("DATA_SOURCE_NAME")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	tagRepo, err := repository.NewTagPostgresRepository(db)
	if err != nil {
		log.Fatal(err)
		return
	}

	//init tag manager
	tagTgManager, err := manager.NewTagManager(tagRepo)
	if err != nil {
		log.Fatal(err)
		return
	}

	//init telegram controler
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	debug := os.Getenv("DEBUG")
	telegramControler, err := controlers.NewTelegramControler(token, host, port, debug, *tagTgManager)
	if err != nil {
		log.Fatal(err)
		return
	}

	telegramControler.StartListen()
}
