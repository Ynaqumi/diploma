package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"main/pkg/status"
	"main/pkg/structs"
	"net/http"
)

var (
	smsChan       = make(chan [][]structs.SMSData)
	mmsChan       = make(chan [][]structs.MMSData)
	voiceCallChan = make(chan []structs.VoiceCallData)
	emailChan     = make(chan map[string][][]structs.EmailData)
	billingChan   = make(chan structs.BillingData)
	supportChan   = make(chan []int)
	incidentsChan = make(chan []structs.IncidentData)
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleConnection).Methods("GET", "OPTIONS")
	http.ListenAndServe("127.0.0.1:8282", router)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	result := getResultData()
	byteResult, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(byteResult)
}

func getResultData() structs.ResultT {

	result := structs.ResultT{false, structs.ResultSetT{}, "Error on collect data"}

	go status.GetSms(smsChan)
	go status.GetMms(mmsChan)
	go status.GetVoice(voiceCallChan)
	go status.GetEmail(emailChan)
	go status.GetBilling(billingChan)
	go status.GetSupport(supportChan)
	go status.GetIncidents(incidentsChan)

	if incidentsChan != nil {

		smsData := <-smsChan
		if len(smsData) == 0 {
			return result
		}

		mmsData := <-mmsChan
		if len(mmsData) == 0 {
			return result
		}

		voiceData := <-voiceCallChan
		if len(voiceData) == 0 {
			return result
		}

		emailData := <-emailChan
		if len(emailData) == 0 {
			return result
		}

		checkBillingData := structs.BillingData{}
		billingData := <-billingChan
		if billingData == checkBillingData {
			return result
		}

		supportData := <-supportChan
		if len(supportData) == 0 {
			return result
		}

		incidentsData := <-incidentsChan
		if len(incidentsData) == 0 {
			return result
		}

		resultSetAns := structs.ResultSetT{smsData, mmsData, voiceData, emailData, billingData, supportData, incidentsData}

		result = structs.ResultT{true, resultSetAns, ""}
	}
	return result
}
