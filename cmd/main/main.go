package main

import (
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Text == "batman" {
				sendBatman(bot, update)
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			if update.Message.Text == "hello" {
				msg.Text = "fuck u"
			} else if update.Message.Text == "fuck u to" {
				msg.Text = "fuck u again"
			} else {
				msg.Text = "я твоя не понимать"
			}

			retMsg, retErr := bot.Send(msg)

			log.Println("sent err:", retErr)
			log.Println("sent msg:", retMsg)
		}
	}
}

var batmanBytes []byte

func sendBatman(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if batmanBytes == nil {
		f, err := ioutil.ReadFile("./data/img/batman.png")
		if err != nil {
			log.Fatalln(err)
		}

		batmanBytes = f
	}

	retMsg, retErr := bot.Send(tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileBytes{
		Name:  "я бэтмен",
		Bytes: batmanBytes,
	}))

	log.Println("sent err:", retErr)
	log.Println("sent msg:", retMsg)
}
