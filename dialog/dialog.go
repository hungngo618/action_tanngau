package dialog

import (
	"github.com/michlabs/fbbot"
)
var ftel *fbbot.Dialog

func New() *fbbot.Dialog {
	d := fbbot.NewDialog()

	var selectbot SelectBot
	var welcome Welcome
	var faq FAQ
	var silence Silence
	var err Error
	var goodbye Goodbye

	d.SetBeginStep(selectbot)
	d.SetEndStep(goodbye)

	d.AddTransition(SelectBotEvent, selectbot)
	d.AddTransition(GoWelcomeEvent, welcome)
	d.AddTransition(GoFAQEvent, faq)
	d.AddTransition(ErrorEvent, err)
	d.AddTransition(GoSlienceEvent, silence)
	d.AddTransition(GoodbyeEvent, goodbye)

	d.PreHandleMessageHook = PreHandlerMessageHook

	ftel = d
	return d
}