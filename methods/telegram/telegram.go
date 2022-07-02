package telegram

import (
	"log"
	"strconv"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Tgconf struct{
	token string
}

type TelegramWorker struct{
	bot *tg.BotAPI
}

func NewTelegramWorker(token string) *TelegramWorker{
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false
//	gin.SetMode(gin.ReleaseMode)

	return &TelegramWorker{bot:bot}
}

func (tw *TelegramWorker) Loop(){
	u := tg.NewUpdate(0)
	u.Timeout = 60
	updates := tw.bot.GetUpdatesChan(u)
	for update := range updates{
		if update.Message != nil{
			log.Printf("[id:%d][%s] %s\n", update.Message.From.ID, update.Message.From.UserName, update.Message.Text)
			msg := tg.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			tw.bot.Send(msg)
		}
	}
}

func (tw *TelegramWorker)SendTgMsg(chatID, data string) error{
	id, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil{
		return err
	}
	msg := tg.NewMessage(id, data)
	if _, err := tw.bot.Send(msg); err != nil {
		return err
        }
	return nil
}
