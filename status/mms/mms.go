package mms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/status/check"
	"main/structs"
	"net/http"
	"sort"
)

func GetMms(mmsChan chan [][]structs.MMSData) {
	request, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		log.Println("Не удалось выполнить GET запрос", err)
		mmsChan <- nil
		return
	}

	if request.StatusCode != 200 {
		var mmsData []structs.MMSData
		log.Println("Ошибка статус код MMS не равен 200", mmsData)
		mmsChan <- nil
		return
	} else {
		fmt.Println("Код ответа 200")
	}

	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println("Не удалось прочитать запрос", err)
		mmsChan <- nil
		return
	}

	var unsortedMms []structs.MMSData
	if err := json.Unmarshal(bytes, &unsortedMms); err != nil {
		log.Println("Ошибка unmarshal", err)
		mmsChan <- nil
		return
	}

	for i := 0; i < len(unsortedMms); i++ {
		if check.CountryCheck(unsortedMms[i].Country) && check.ProviderSmsAndMms(unsortedMms[i].Provider) && unsortedMms[i].ResponseTime != "" && unsortedMms[i].Bandwidth != "" {
			unsortedMms = unsortedMms[1 : len(unsortedMms)-1]
		}
	}
	sortedMmsCollection := [][]structs.MMSData{}
	mmsSortedProvider := sortProviderMms(unsortedMms)
	sortedMmsCollection = append(sortedMmsCollection, mmsSortedProvider)
	mmsSortedCountry := sortCountryMms(unsortedMms)
	sortedMmsCollection = append(sortedMmsCollection, mmsSortedCountry)

	mmsChan <- sortedMmsCollection
}

func sortProviderMms(unsortedMms []structs.MMSData) []structs.MMSData {
	mmsSortedProvider := make([]structs.MMSData, len(unsortedMms))
	copy(mmsSortedProvider, unsortedMms)
	sort.Slice(mmsSortedProvider, func(i, j int) bool {
		return mmsSortedProvider[i].Provider < mmsSortedProvider[j].Provider
	})
	return mmsSortedProvider
}

func sortCountryMms(unsortedMms []structs.MMSData) []structs.MMSData {
	mmsSortedCountry := make([]structs.MMSData, len(unsortedMms))
	copy(mmsSortedCountry, unsortedMms)
	sort.Slice(mmsSortedCountry, func(i, j int) bool {
		return mmsSortedCountry[i].Country[:2] < mmsSortedCountry[j].Country[:2]
	})
	return mmsSortedCountry
}
