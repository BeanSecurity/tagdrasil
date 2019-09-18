package main

import (
	"tagdrasil/controlers"
	"tagdrasil/repository"
	"tagdrasil/manager"
	"os"
)

func main() {
	tagRepo := repository.NewTagPostgresRepository("")
	// var tagTgManager *manager.TagManager = manager.NewTagManager(tagRepo)
	tagTgManager := manager.NewTagManager(tagRepo)

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	debug := os.Getenv("DEBUG")
	// dburl := os.Getenv("DATABASE_URL")
	dsn := os.Getenv("DATA_SOURCE_NAME")
	telegramControler := controlers.NewTelegramControler(token, host, port, debug, dsn, *tagTgManager)

	telegramControler.StartListen()
}
