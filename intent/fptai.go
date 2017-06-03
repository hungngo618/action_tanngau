package intent

import (
	"fmt"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	fptai "github.com/fpt-corp/fptai-sdk-go"
)

type FPTAI struct {
	app *fptai.Client
}

func NewFPTAI(config map[string]string) (Engine, error) {
	applicationToken := config["application_token"]
	if applicationToken == "" {
		return nil, fmt.Errorf("FPT.AI application_token must not be empty")
	}

	client := fptai.NewClient(applicationToken)

	var f FPTAI
	f.app = client

	return &f, nil
}

func (f *FPTAI) Detect(text string) string {
	resp, err := f.app.RecognizeIntents(strings.ToLower(text))
	if err != nil {
		log.Errorf("FPT.AI failed to recognize intent. Error: %s. Text: %s\n", err.Error(), text)
		return ""
	}

	if len(resp.Intents) > 0 {
		confidence := resp.Intents[0].Confidence

		if confidence < 0.7 {
			return ""
		}

		return resp.Intents[0].Name
	}

	return ""
}

func (f *FPTAI) AddUtterance(intent, utterance string) error {
	if err := f.app.CreateUtterances(intent, []string{utterance}); err != nil {
		log.Errorf("FPT.AI failed to add new utterance for intent. Error: %s. Intent: %s. Utterance: %s\n", err.Error(), intent, utterance)
		return err
	}
	return nil
}

func (f *FPTAI) AddIntentUtterances(ius []IntentUtterance) error {
	intents, err := f.app.GetIntents()
	if err != nil {
		log.Error("FPTAI failed to get intents: ", err)
		return err
	}

	existingIntents := make(map[string]bool)
	for _, i := range intents {
		existingIntents[i.Name] = true
	}

	var errs []error
	for idx, iu := range ius {
		if iu.Utterance == ""{
			continue
		}

		if !existingIntents[iu.Intent] {
			i, err := f.app.CreateIntent(iu.Intent, iu.Intent)
			if err != nil {
				log.Error("FPTAI failed to create new intent: ", err)
				errs = append(errs, err)
				continue
			}
			existingIntents[i.Name] = true
		}

		for i:=0; i<3; i++ {
			err := f.app.CreateUtterances(iu.Intent, []string{strings.ToLower(iu.Utterance)})
			if err != nil {
				time.Sleep(2 * time.Second)
				continue
			}
			break
		}

		if err != nil {
			log.Error("Failed in utt %d - %s",idx, iu.Utterance)
		}
	}

	if len(errs) > 0 {
		log.Errorf("Erros:\n")
		for i, e := range errs {
			log.Errorf("%d\t%+v\n", i, e)
		}
	}
	
	return nil
}

func (f *FPTAI) DeleteAllIntents() error {
	intents, err := f.app.GetIntents()
	if err != nil {
		log.Error("FPT AI failed to get all intents: ", err.Error())
		return err
	}

	for _, intent := range(intents) {
		err := f.app.DeleteIntent(intent.Name)
		if err != nil {
			log.Errorf("Failed to delete intent: %s %v", err.Error(), intent)
			return err
		}
	}

	log.Info("Deleted success")

	return nil
}

func (f *FPTAI) Train() error {
	return f.app.TrainIntent()
}