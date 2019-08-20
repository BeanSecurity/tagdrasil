package tagdrasil

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
	manager "tagdrasil/manager"
)

type TelegramControler struct {
	TagManager manager.TagManager
	port       string
	bot        *tgbotapi.BotAPI
}

func NewTelegramControler(manager manager.TagManager) *TelegramControler {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	// dburl := os.Getenv("DATABASE_URL")
	// dsn := os.Getenv("DATA_SOURCE_NAME")
	url := host + ":443/" + token

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	return &TelegramControler{TagManager: manager, port: port, bot: bot}
}

// func (t *TelegramControler) UpdateProcess() {
// }

// func (t *TelegramControler) commandFacade(command string) {
// }
func (t *TelegramControler) StartListen() {
	updates := t.bot.ListenForWebhook("/" + t.bot.Token)
	go http.ListenAndServe(":"+t.port, nil)

	for update := range updates {
		log.Printf("%+v\n", update)
		t.processTelegramUpdate(update)
	}
}

func (t *TelegramControler) processTelegramUpdate(update tgbotapi.Update) {
	if update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		// commandSelect:
		switch update.Message.Command() {
		// case "start":
		// case "add":
		// case "tags":
		default:
			msg.Text = "I dont know that command"
		}
	}

	if update.ChannelPost != nil {
		log.Printf("%+v\n", update.ChannelPost)
		log.Printf("%+v\n", update.ChannelPost.Chat)
		_, err := t.bot.Send(tgbotapi.NewEditMessageText(update.ChannelPost.Chat.ID, update.ChannelPost.MessageID, "AAAAAAAA"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
