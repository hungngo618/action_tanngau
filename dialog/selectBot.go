package dialog

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/fpt-corp/util"
	"github.com/michlabs/fbbot"
)

type SelectBot struct {
	fbbot.BaseStep
}

func (s SelectBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	mess := strings.ToLower(util.RemoveAccent(msg.Text))
	logrus.Infof("mess %s compare %s %s \n", mess, strings.ToLower(util.RemoveAccent(T("select_human"))), strings.ToLower(util.RemoveAccent(T("select_bot"))))
	if mess != strings.ToLower(util.RemoveAccent(T("select_human"))) && mess != strings.ToLower(util.RemoveAccent(T("select_bot"))){
		if bot.STMemory.For(msg.Sender.ID).Get("question") == "" && !strings.HasPrefix(msg.Text, "/"){
			bot.STMemory.For(msg.Sender.ID).Set("question", msg.Text)
		}
	}
	var items  []fbbot.QuickRepliesItem
	items = append(items, fbbot.QuickRepliesItem{
		ContentType: "Text",
		Title:       T("select_human"),
		Payload:     "yes",
		// ImageURL:    T("select_human_icon"),
	})
	items = append(items, fbbot.QuickRepliesItem{
		ContentType: "Text",
		Title:       T("select_bot"),
		Payload:     "no",
		// ImageURL:    T("select_bot_icon"),
	})

	bot.SendText(msg.Sender, util.Personalize(T("select_bot_title"), &msg.Sender))
	botSelect := new(fbbot.QuickRepliesMessage)
	botSelect.Text = util.Personalize(T("select_bot_options"), &msg.Sender)
	botSelect.Items = items

	bot.Send(msg.Sender, botSelect)

	return StayEvent
}


func (s SelectBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	}

	switch strings.ToLower(util.RemoveAccent(msg.Text)) {
	case strings.ToLower(util.RemoveAccent(T("select_human"))):
		return GoSlienceEvent
	case strings.ToLower(util.RemoveAccent(T("select_bot"))):
		bot.STMemory.For(msg.Sender.ID).Set("previous_step", "select")
		return GoWelcomeEvent
	}
	return s.Enter(bot, msg)
}