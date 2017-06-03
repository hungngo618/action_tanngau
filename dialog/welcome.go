package dialog

import (
	"github.com/fpt-corp/util"
	"github.com/michlabs/fbbot"
)

type Welcome struct {
	fbbot.BaseStep
}

func (s Welcome) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	isFirstConversation := bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation")

	if isFirstConversation == "false" {
		bot.TypingOn(msg.Sender)
		bot.SendText(msg.Sender, util.Personalize(T("welcome_second_time"), &msg.Sender))
		bot.STMemory.For(msg.Sender.ID).Delete("question")
		return GoFAQEvent
	} else {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "false")
		bot.TypingOn(msg.Sender)

		bot.SendText(msg.Sender, util.Personalize(T("welcome_first_time"), &msg.Sender))
		bot.STMemory.For(msg.Sender.ID).Delete("question")
		return GoFAQEvent
	}
}
