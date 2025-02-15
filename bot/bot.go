package bot

import (
	"fmt"
	"log"
	"republish/observer"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	TOCKEN = "7313692442:AAG8jCBlrC4NOZwJQKxiglFhxsdbTHS0BYA"
)

type Bot struct {
	Tgbot *tgbotapi.BotAPI
	obs *observer.Observer
}

func NewBot() (*Bot, error) {
	tgbot, err := tgbotapi.NewBotAPI(TOCKEN)
	if err != nil {
		return nil, fmt.Errorf("Can not create bot: %v", err)
	}

	tgbot.Debug = true
	bot := &Bot{
		Tgbot: tgbot,
		obs: observer.NewObserver(),
	}

	return bot, nil
}

func (bot *Bot) StartBot() {
	log.Printf("Authorized on account %s", bot.Tgbot.Self.UserName)

	go func() {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := bot.Tgbot.GetUpdatesChan(u)

		for update := range updates {
			if update.ChannelPost != nil {
				bot.obs.PushMessageToRouting(update, bot.Tgbot)
			}
		}
	}()
}

func (b *Bot) SendHelloWorld(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage((*message).Chat.ID, "Hello, world!")
	b.Tgbot.Send(msg)
	return nil
}