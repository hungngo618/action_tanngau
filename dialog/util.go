package dialog

import (
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/fpt-corp/action-tanngau/db"
	"github.com/fpt-corp/action-tanngau/intent"
	"github.com/fpt-corp/util"
	"github.com/michlabs/fbbot"
)

func HandleQuestion(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	question := bot.STMemory.For(msg.Sender.ID).Get("question")
	if question == "" {
		question = msg.Text
	}
	log.Infof("Question = %s", question)

	bot.STMemory.For(msg.Sender.ID).Delete("question")
	i := intent.Detect(question)
	switch i {
	case "":
		log.Debug("There's no matching sentence for: ", msg.Text)
		bot.SendText(msg.Sender, util.Personalize(T("dontunderstand"), &msg.Sender))
		return StayEvent
	case "goodbye":
		return GoodbyeEvent
	default:
		a := db.GetAnswerFor(i)
		if a == "" {
			bot.SendText(msg.Sender, util.Personalize(T("no_answer"), &msg.Sender))
			return NoAnswerEvent
		}

		var arr []string
		a = strings.TrimSpace(a)
		arr = strings.SplitN(a, "\n\n", -1)
		for _, v := range arr {
			bot.TypingOn(msg.Sender)
			if v != "" {
				bot.TypingOn(msg.Sender)
				time.Sleep(2 * time.Second)
				util.SendTextWithImages(bot, msg.Sender, util.Personalize(v, &msg.Sender))
			}
		}
	}
	return StayEvent
}
