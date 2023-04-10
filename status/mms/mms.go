package mms

import (
	"diploma2/status/check"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func Mms() (mms []MMSData) {
	request, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		log.Printf("Не удалось выполнить GET запрос по MMS. Код ответа %v. Ошибка %v \n", request.StatusCode, err)
	} else {
		fmt.Printf("GET запрос по MMS выполнен. Код ответа %v \n", request.StatusCode)
	}

	mmsData := []MMSData{}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println("Не удалось прочитать Get-запрос", err)
	}

	if err := json.Unmarshal(body, &mmsData); err != nil {
		log.Println("Ошибка unmarshal", err)
	}

	for _, elem := range mmsData {
		if check.CountryCheck(elem.Country) && check.ProviderSmsAndMMSCheck(elem.Provider) {
			mms = append(mms, MMSData{Country: elem.Country + ";", Provider: elem.Provider + ";", Bandwidth: elem.Bandwidth + ";", ResponseTime: elem.ResponseTime})
		}
	}
	return
}
