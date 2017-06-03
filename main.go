package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/fpt-corp/action-tanngau/config"
	"github.com/fpt-corp/action-tanngau/db"
	"github.com/fpt-corp/action-tanngau/dialog"
	"github.com/fpt-corp/action-tanngau/intent"
	"github.com/fpt-corp/action-tanngau/ui"
	"github.com/michlabs/fbbot"
)

var ftel *fbbot.Dialog

func init() {
	if err := config.LoadFromEnv(); err != nil {
		log.Fatal("failed to load configuration: ", err)
	}

	if err := db.Init(&config.DB); err != nil {
		log.Fatal("failed to connect to db: ", err)
	}

	if err := intent.Init("fptai", config.Bot.FPTAI); err != nil {
	//if err := intent.Init("wit", config.Bot.Wit); err != nil {
		log.Fatal("failed to init intent package: ", err)
	}
	
	dialog.Init(config.Bot.LanguageFile)
}

func main() {



	switch config.Bot.LogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	logWriter, err := os.OpenFile(config.Bot.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logWriter)

	// debugLogWriter, err := os.OpenFile(config.Bot.DebugFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer debugLogWriter.Close()
	// debug.Init(bot, debugLogWriter)

	ui.Init(&config.UI)
	ui.Run()

}
