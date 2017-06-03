package config

import(
	"github.com/kelseyhightower/envconfig"
)

var Bot BotConfig
var UI UIConfig
var DB DBConfig

type BotConfig struct {
	Port            int    `default:"1203"`
	VerifyToken     string `required:"true"`
	PageAccessToken string `required:"true"`
	LogFile string `default:"bot.log"`
	LogLevel        string `default:"info"`
	Wit        map[string]string `required:"true"`
	FPTAI	map[string]string `required:"true"`
	LanguageFile        string `required:"true"`
	DebugFile	string `required:"true"`
	Hex string `required:"true"`
	ConversationTimeout float64 `required:"true"`
	SilenceDuration float64 `required:"true"`
}

type UIConfig struct {
	Port       int    `required:"true"`
	Username   string `required:"true"`
	Userpass   string `required:"true"`
	StaticDir  string `required:"true"`
	BotName	   string `required:"true"`
}

type DBConfig struct {
	Host     string `required:"true"`
	Port	 string `default:"3306"`
	Name     string `required:"true"`
	User     string `required:"true"`
	Pass string `required:"true"`
}

func LoadFromEnv() error {
	if err := envconfig.Process("ftel", &Bot); err != nil {
		return err
	}

	if err := envconfig.Process("ui", &UI); err != nil {
		return err
	}

	if err := envconfig.Process("db", &DB); err != nil {
		return err
	}

	return nil
}