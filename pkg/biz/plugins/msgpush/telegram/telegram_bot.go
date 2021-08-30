package telegram

import (
	"github.com/d0zingcat/seafood/pkg/httpclient"
	"github.com/d0zingcat/seafood/pkg/utils"
)

type TelegramBotEnt struct {
	Token string
}

type TelegramBot interface {
	SendMessage(Message)
}

func (t *TelegramBotEnt) SendMessage(m Message) {
	u := utils.Conf.TelegramBotAPIURL + utils.Conf.TelegramBotToken + "/sendMessage"
	httpclient.BuildRequst().Post(u)
}

type Message struct {
	ChatID string `schama:"chat_id"`
	Text   string `schema:"text"`
}

func NewBot(token string) (bot TelegramBot) {
	telegramEnt := &TelegramBotEnt{
		Token: token,
	}
	bot = telegramEnt
	return
}
