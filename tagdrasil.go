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
	dburl := os.Getenv("DATABASE_URL")
	log.Printf("dburl: %s\n", dburl)
	dsn := os.Getenv("DATA_SOURCE_NAME")
	log.Printf("dsn: %s\n", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	tagRepo, err := repository.NewTagPostgresRepository(db)
	if err != nil {
		log.Println(err)
		return
	}

	//init tag manager
	tagTgManager, err := manager.NewTagManager(tagRepo)
	if err != nil {
		log.Println(err)
		return
	}

	//init telegram controler
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	debug := os.Getenv("DEBUG")
	telegramControler, err := controlers.NewTelegramControler(token, host, port, debug, *tagTgManager)
	if err != nil {
		log.Println(err)
		return
	}

	telegramControler.StartListen()
}
