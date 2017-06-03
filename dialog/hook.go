package dialog

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/fpt-corp/action-tanngau/config"
	"github.com/fpt-corp/util"
	"github.com/michlabs/fbbot"
)

func checkTiming(bot *fbbot.Bot, msg *fbbot.Message) bool {
	lastEcho := bot.STMemory.For(msg.Sender.ID).Get("lastEcho")

	if lastEcho != "" && !TimeExpired(lastEcho, config.Bot.SilenceDuration) {
		log.Debugf("Bot do nothing since staff is chatting, lastEcho = %s", util.ToTime(lastEcho).Format("15:04:05"))
		return true
	}

	lastActiveTime := bot.STMemory.For(msg.Sender.ID).Get("lastMessage")

	if TimeExpired(lastActiveTime, config.Bot.ConversationTimeout) {
		log.Debug("Dialog is expired, move to Welcome state")
		ftel.Move(msg, SelectBot{})
		return true
	}

	return false
}

func AdHocPostBack(bot *fbbot.Bot, msg *fbbot.Message) bool {
	log.Infof("\n msg = %s \n", msg.Text)
	log.Infof("\n len = %d %d %d %d \n", len(msg.Audios), len(msg.Files), len(msg.Images), len(msg.Videos))
	if len(msg.Audios) > 0 || len(msg.Files) > 0 || len(msg.Images) > 0 || len(msg.Videos) > 0 {
		bot.SendText(msg.Sender, util.Personalize(T("dontunderstand"), &msg.Sender))
		bot.STMemory.For(msg.Sender.ID).Set("previous_step", "Emoji")
		ftel.Move(msg, FAQ{})
		return true
	}

	emo := GetEmojiResponse(msg.Text)
	if emo != "" {
		bot.SendText(msg.Sender, emo)
		bot.STMemory.For(msg.Sender.ID).Set("previous_step", "Emoji")
		ftel.Move(msg, FAQ{})
		return true
	}

	switch strings.ToLower(msg.Text) {
	case "/restart":
		bot.STMemory.For(msg.Sender.ID).Set("previous_step", "hook")
		ftel.Move(msg, Goodbye{})
		bot.STMemory.For(msg.Sender.ID).Delete("lastEcho")
		bot.STMemory.For(msg.Sender.ID).Delete("lastMessage")
		return true
	case "/bot":
		bot.STMemory.For(msg.Sender.ID).Set("previous_step", "hook")
		ftel.Move(msg, Welcome{})
		bot.STMemory.For(msg.Sender.ID).Delete("lastEcho")
		bot.STMemory.For(msg.Sender.ID).Delete("lastMessage")
		return true

	}

	return false
}

func PreHandlerMessageHook(bot *fbbot.Bot, msg *fbbot.Message) bool{
	return AdHocPostBack(bot, msg) || checkTiming(bot, msg)
}