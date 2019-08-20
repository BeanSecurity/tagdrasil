package main

import (
	controler "tagdrasil/controlers"
	manager "tagdrasil/manager"
	repository "tagdrasil/repositories"
)

func main() {
	tagRepo := repository.NewTagPostgresRepository()
	var tagTgManager manager.TagManager = manager.NewTagTelegramManager(tagRepo)
	var telegramControler *controler.TelegramControler = controler.NewTelegramControler(tagTgManager)

	telegramControler.StartListen()
}
