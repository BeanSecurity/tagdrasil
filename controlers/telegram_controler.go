package controlers

import (
	"errors"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/xlab/treeprint"
	"log"
	"net/http"
	"regexp"
	"strings"
	"tagdrasil/manager"
	"tagdrasil/models"
)

type TelegramControler struct {
	TagManager manager.TagManager
	port       string
	bot        *tgbotapi.BotAPI
}

func NewTelegramControler(token, host, port, debug string, manager manager.TagManager) (*TelegramControler, error) {
	url := host + ":443/" + token

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if debug == "1" {
		bot.Debug = true
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	return &TelegramControler{TagManager: manager, port: port, bot: bot}, nil
}

func (t *TelegramControler) StartListen() {
	updates := t.bot.ListenForWebhook("/" + t.bot.Token)
	go http.ListenAndServe(":"+t.port, nil)

	for update := range updates {
		log.Printf("%+v\n", update)
		t.processTelegramUpdate(update)
	}
}

func (t *TelegramControler) processTelegramUpdate(upd tgbotapi.Update) {
	if upd.Message.IsCommand() {
		var err error
		msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "huh")
		// commandSelect:
		switch upd.Message.Command() {
		case "start":
			msg.Text = "hi"
			err := t.userStart(upd)
			if err != nil {
				log.Fatal(err)
				msg.Text = "sorry, еггог"
			}
		case "add":
			msg.Text = "tag added"
			err = t.addTags(upd)
			if err != nil {
				log.Fatal(err)
				msg.Text = "sorry, еггог"
			}
		// case "tags":
		default:
			msg.Text = "I dont know that command"
		}
		_, err = t.bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	} else if upd.Message != nil {
		log.Printf("%+v\n", upd.Message.Text)
		log.Printf("%+v\n", upd.Message.From)
		var newText string
		var tagHeader models.TagNode

		tags, err := t.ParseTags(upd.Message.Text)
		if err != nil {
			log.Fatal(err)
			_, err = t.bot.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "sorry, error"))
		}

		user := models.User{
			ID:        upd.Message.From.ID,
			FirstName: upd.Message.From.FirstName,
			LastName:  upd.Message.From.LastName,
			UserName:  upd.Message.From.UserName,
		}

		tagHeader, err = t.TagManager.GetTagHeader(tags, user)
		if err != nil {
			log.Fatal(err)
			_, err = t.bot.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "sorry, error"))
		}

		newText, err = t.TagIntoText(tagHeader)
		if err != nil {
			log.Fatal(err)
			_, err = t.bot.Send(tgbotapi.NewMessage(upd.Message.Chat.ID, "sorry, error"))
		}

		_, err = t.bot.Send(tgbotapi.NewMessage(
			upd.Message.Chat.ID,
			newText))
		if err != nil {
			log.Fatal(err)
		}
	}

	// if upd.ChannelPost != nil {
	// log.Printf("%+v\n", upd.ChannelPost)
	// log.Printf("%+v\n", upd.ChannelPost.From)
	// var newText string

	// tags := t.ParseTags(upd.ChannelPost.Text)
	// user := models.User{}
	// tagHeader := t.TagManager.GetTagHeader(tags, upd.Message.From.ID)
	// _, err = t.bot.Send(tgbotapi.NewEditMessageText(upd.ChannelPost.From.ID, upd.ChannelPost.MessageID, newText))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// }
}

// addTags ...
func (t *TelegramControler) addTags(upd tgbotapi.Update) error {
	args := strings.Fields(upd.Message.CommandArguments())
	switch len(args) {
	case 1:
		user := mapTgUserToModelUser(*upd.Message.From)
		err := t.TagManager.AddRootTag(args[0], user)
		return err
	case 2:
		user := mapTgUserToModelUser(*upd.Message.From)
		err := t.TagManager.AddLeafTag(args[0], args[1], user)
		return err
	}
	return errors.New("too many arguments")
}

func (t *TelegramControler) userStart(upd tgbotapi.Update) error {
	tgUser := upd.Message.From
	user := mapTgUserToModelUser(*tgUser)
	err := t.TagManager.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (t *TelegramControler) ParseTags(text string) ([]models.TagNode, error) {
	re := regexp.MustCompile(`#(\w*[A-Za-z_]+\w*)`)
	tags := re.FindAllStringSubmatch(text, -1)
	tagNodes := make([]models.TagNode, 0)
	for _, tag := range tags {
		tagNodes = append(tagNodes, models.TagNode{Name: tag[1]})
	}
	return tagNodes, nil
}

func (t *TelegramControler) TagIntoText(tag models.TagNode) (string, error) {
	if t.IsLine(tag) {
		s, err := t.TagLineIntoText(tag)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		return s, nil
	} else {
		tree := treeprint.New()
		tree.SetValue("#" + tag.Name)
		for _, tag := range tag.ChildTags {
			err := t.TagTreeIntoText(tree, tag)
			if err != nil {
				log.Fatal(err)
				return "", err
			}
		}
		return tree.String(), nil
	}
}

func (t *TelegramControler) IsLine(tag models.TagNode) bool {
	switch len(tag.ChildTags) {
	case 0:
		return true
	case 1:
		return t.IsLine(tag.ChildTags[0])
	default:
		return false
	}
}

func (t *TelegramControler) TagLineIntoText(tag models.TagNode) (string, error) {
	switch len(tag.ChildTags) {
	case 0:
		return "#" + tag.Name, nil
	case 1:
		s, err := t.TagLineIntoText(tag.ChildTags[0])
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		return "#" + tag.Name + " -> " + s, nil
	default:
		err := errors.New("Tag is not the tag line")
		log.Fatal(err)
		return "", err
	}

}

func (t *TelegramControler) TagTreeIntoText(tree treeprint.Tree, tag models.TagNode) error {
	var branch treeprint.Tree

	if len(tag.ChildTags) == 0 {
		tree.AddNode("#" + tag.Name)
	} else {
		branch = tree.AddBranch("#" + tag.Name)
		for _, currentNode := range tag.ChildTags {
			err := t.TagTreeIntoText(branch, currentNode)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// mapTgUserToModelUser ...
func mapTgUserToModelUser(tgUser tgbotapi.User) models.User {
	return models.User{
		ID:        tgUser.ID,
		UserName:  tgUser.UserName,
		FirstName: tgUser.FirstName,
		LastName:  tgUser.LastName,
	}
}
