package ui

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/fpt-corp/action-tanngau/config"
	"github.com/gorilla/mux"
	//"github.com/justinas/alice"
	"github.com/nguyendangminh/gotpl"
)

var tplHelper *gotpl.Helper
var HistoriesPerPage int = 50
var isTraining bool = false
var staticPath string

func Init(cfg *config.UIConfig) error {
	var err error
	tplHelper, err = gotpl.New(fmt.Sprintf("%s/view", config.UI.StaticDir))
	if err != nil {
		return err
	}

	return nil
}

func TrainHandler(w http.ResponseWriter, r *http.Request) {
	if (!isTraining) {
		isTraining = true
		go FAQTrain(w, r)
	}
}

func CheckTrainStatus(w http.ResponseWriter, r *http.Request) {
	if (isTraining) {
		fmt.Fprintln(w, "Training")
	} else {
		fmt.Fprintln(w, "Done")
	}
}

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	FileName := staticPath + "/ftel_faq.xlsx"
	log.Info("Path is: ", FileName)
	err := Export(FileName)
	if err != nil {
		log.Error("Fail to export excel: ", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Info("redirect")
	http.Redirect(w, r, "/resources/ftel_faq.xlsx", http.StatusFound)

	go DeleteFile(FileName, 60)

	return
}

func Run() {
	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler)

	r.HandleFunc("/login", Login)
	r.HandleFunc("/logout", Logout)

	r.HandleFunc("/faq", ListFAQsHandler).Methods("GET")
	r.HandleFunc("/newfaq", NewFAQHandler).Methods("GET", "POST")
	r.HandleFunc("/faq/{id:[0-9]+}", DeleteFAQHandler).Methods("DELETE")

	r.HandleFunc("/editview", FAQEditViewRender)
	r.HandleFunc("/editfaq", FAQEditHandler)

	r.HandleFunc("/deletequestion", DeleteQuestionHandler).Methods("POST")

	r.HandleFunc("/search", SearchFAQByIntentHandler).Methods("POST", "GET")
	r.HandleFunc("/qna", SearchAnswerByQuestion).Methods("POST", "GET")

	r.HandleFunc("/train", TrainHandler)

	r.HandleFunc("/trainstatus", CheckTrainStatus)

	r.HandleFunc("/help", HelpHandler)
	r.HandleFunc("/export", ExportHandler)

	r.HandleFunc("/unsafe", UnsafeHandler)

	staticPath = fmt.Sprintf("%s/resources", config.UI.StaticDir)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir(staticPath))))

	//chain := alice.New(author).Then(r)
	http.Handle("/", r)

	log.Info("FAQ UI is running at: ", config.UI.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.UI.Port), nil))
}
