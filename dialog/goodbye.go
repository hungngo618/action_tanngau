package dialog

import (
	"github.com/fpt-corp/action-tanngau/db"
	"github.com/fpt-corp/util"
	"github.com/michlabs/fbbot"
)

type Goodbye struct {
	fbbot.BaseStep
}

func (s Goodbye) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	reply := db.GetAnswerFor("goodbye")
	if reply == "" {
		reply = T("goodbye")
	}

	reply = util.Personalize(reply, &msg.Sender)
	bot.SendText(msg.Sender, reply)
	
	return StayEvent
}
