package server

import (
	"diploma/internal/data/billing"
	"diploma/internal/data/email"
	"diploma/internal/data/incidents"
	"diploma/internal/data/mms"
	"diploma/internal/data/sms"
	"diploma/internal/data/support"
	"diploma/internal/data/voice"
	"diploma/internal/structs"
	"encoding/json"
	"log"
	"net/http"
)

import (
	"github.com/gorilla/mux"
)

func Start() {
	server := &http.Server{
		Addr: "localhost:8282",
	}

	router := mux.NewRouter()
	router.HandleFunc("/api", handleConnection)
	server.Handler = router

	log.Fatal(server.ListenAndServe())
}

func handleConnection(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	resultT := getResultT()
	byteResultT, err := json.Marshal(resultT)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Write(byteResultT)
}

func getResultT() (resultT structs.ResultT) {
	result := []structs.ResultT{}
	resultSet := getResultData(result)
	resultT = structs.ResultT{Status: true, Data: resultSet}
	return
}

func getResultData(resultT []structs.ResultT) (resultSet structs.ResultSetT) {
	sms, errSms := sms.Sms()
	if errSms != "" {
		resultT = append(resultT, structs.ResultT{Status: false, Data: structs.ResultSetT{}, Error: errSms})
		return
	}
	mms, errMms := mms.Mms()
	if errMms != "" {
		resultT = append(resultT, structs.ResultT{Status: false, Data: structs.ResultSetT{}, Error: errMms})
		return
	}
	voice, errVoice := voice.VoiceCall()
	if errVoice != "" {
		resultT = append(resultT, structs.ResultT{Status: false, Data: structs.ResultSetT{}, Error: errVoice})
		return
	}
	email, errEmail := email.Email()
	if errEmail != "" {
		resultT = append(resultT, structs.ResultT{Status: false, Data: structs.ResultSetT{}, Error: errEmail})
		return
	}
	billing, errBilling := billing.Billing()
	if errBilling != "" {
		resultT = append(resultT, structs.ResultT{Status: false, Data: structs.ResultSetT{}, Error: errBilling})
		return
	}
	support, errSupport := support.Support()
	if errSupport != "" {
		resultT = append(resultT, structs.ResultT{Status: false, Data: structs.ResultSetT{}, Error: errSupport})
		return
	}
	incidents, errIncidents := incidents.Incidents()
	if errIncidents != "" {
		resultT = append(resultT, structs.ResultT{Status: false, Data: structs.ResultSetT{}, Error: errIncidents})
		return
	}

	resultSet = structs.ResultSetT{SMS: sms, MMS: mms, VoiceCall: voice, Email: email, Billing: billing, Support: support, Incidents: incidents}

	return
}
