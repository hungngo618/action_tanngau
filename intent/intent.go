package intent
import (
	"strings"
	"github.com/fpt-corp/qna/db"
)

type IntentUtterance struct {
	Intent string `json:"intent"`
	Utterance string `json:"utterance"`
}

// Detect returns intent of a text. 
func Detect(text string) string {
	i := engine.Detect(strings.ToLower(text))


	return i
}

func AddUtterance(intent, utterance string) error {
	return engine.AddUtterance(intent, utterance)
}

func AddIntentUtterances(ius []IntentUtterance) error {
	return engine.AddIntentUtterances(ius)
}
func DeleteAllIntents() error {
	return engine.DeleteAllIntents()
}

func Train() error {
	return engine.Train()
}

func GetAll() ([]string, error) {
	return db.GetAllIntents()
}