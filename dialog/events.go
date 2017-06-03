package dialog

import (
	"github.com/michlabs/fbbot"
)

const SelectBotEvent fbbot.Event    = "select bot"
const StayEvent fbbot.Event         = ""
const GoWelcomeEvent fbbot.Event    = "go welcome"
const GoFAQEvent fbbot.Event        = "go to FAQ"
const ErrorEvent fbbot.Event        = "error"
const GoodbyeEvent fbbot.Event      = "goodbye"
const GoSlienceEvent fbbot.Event    = "go slience"
const StaffSupportEvent fbbot.Event = "staff support"
const NoAnswerEvent fbbot.Event		= "No answer"
const StaffRegisterEvent fbbot.Event= "Staff register"