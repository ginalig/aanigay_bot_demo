package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMessageText := ""
	products := c.productService.List()

	for _, p := range products {
		outputMessageText += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessageText)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["list"] = (*Commander).List
}
