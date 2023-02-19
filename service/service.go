package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"main/status/billing"
	"main/status/email"
	"main/status/incidents"
	"main/status/mms"
	"main/status/sms"
	"main/status/support"
	"main/status/voiceCall"
	"main/structs"
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
	router.HandleFunc("/config", handleConnection).Methods("GET", "OPTIONS")
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

	result := structs.ResultT{Status: false, Data: structs.ResultSetT{}, Error: "Error on collect data"}

	go sms.GetSms(smsChan)
	go mms.GetMms(mmsChan)
	go voiceCall.GetVoice(voiceCallChan)
	go email.GetEmail(emailChan)
	go billing.GetBilling(billingChan)
	go support.GetSupport(supportChan)
	go incidents.GetIncidents(incidentsChan)

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

		resultSetAns := structs.ResultSetT{SMS: smsData, MMS: mmsData, VoiceCall: voiceData, Email: emailData, Billing: billingData, Support: supportData, Incidents: incidentsData}

		result = structs.ResultT{Status: true, Data: resultSetAns}
	}
	return result
}
