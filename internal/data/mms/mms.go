package mms

import (
	"diploma/internal/check"
	"diploma/internal/structs"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Mms() (mms []structs.MMSData) {
	request, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		log.Printf("Не удалось выполнить GET запрос по MMS. Код ответа %v. Ошибка %v \n", request.StatusCode, err)
	} else {
		fmt.Printf("GET запрос по MMS выполнен. Код ответа %v \n", request.StatusCode)
	}

	mmsData := []structs.MMSData{}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Не удалось прочитать Get-запрос. Ошибка: %v", err)
	}

	if err := json.Unmarshal(body, &mmsData); err != nil {
		log.Printf("Ошибка unmarshal %v", err)
	}

	for _, elem := range mmsData {
		if check.CountryCheck(elem.Country) && check.ProviderSmsAndMMSCheck(elem.Provider) {
			mms = append(mms, structs.MMSData{Country: elem.Country + ";", Provider: elem.Provider + ";", Bandwidth: elem.Bandwidth + ";", ResponseTime: elem.ResponseTime})
		}
	}
	return
}
