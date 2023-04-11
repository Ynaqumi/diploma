package mms

import (
	"diploma/internal/structs"
	"diploma/internal/support_functoins"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
)

func Mms() ([][]structs.MMSData, string) {
	request, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		log.Printf("Не удалось выполнить GET-запрос по MMS. Код ответа %v. Ошибка %v \n", request.StatusCode, err)
		return nil, support_functoins.ErrorToString(err)
	}

	mmsData := []structs.MMSData{}
	unsortedMms := []structs.MMSData{}
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Не удалось прочитать GET-запрос по MMS. Ошибка: %v", err)
		return nil, support_functoins.ErrorToString(err)
	}

	if err := json.Unmarshal(body, &mmsData); err != nil {
		log.Printf("Ошибка unmarshal по MMS: %v", err)
		return nil, support_functoins.ErrorToString(err)
	}

	for _, elem := range mmsData {
		if support_functoins.CountryCheck(elem.Country) && support_functoins.ProviderSmsAndMMSCheck(elem.Provider) {
			unsortedMms = append(unsortedMms, structs.MMSData{Country: elem.Country + ";", Provider: elem.Provider + ";", Bandwidth: elem.Bandwidth + ";", ResponseTime: elem.ResponseTime})
		}
	}

	mmsFullCountry := makeFullCountry(unsortedMms)
	mmsSortedByProvider := sortByProvider(mmsFullCountry)
	mmsSortedByCountry := sortByCountry(mmsFullCountry)
	return [][]structs.MMSData{mmsSortedByProvider, mmsSortedByCountry}, support_functoins.ErrorToString(err)
}

func makeFullCountry(shortnameCountry []structs.MMSData) []structs.MMSData {
	for i := range shortnameCountry {
		if fullCountry, ok := support_functoins.CountriesList[shortnameCountry[i].Country]; ok {
			shortnameCountry[i].Country = fullCountry
		}
	}
	return shortnameCountry
}

func sortByProvider(unsortedMms []structs.MMSData) []structs.MMSData {
	mmsSortedByProvider := make([]structs.MMSData, len(unsortedMms))
	copy(mmsSortedByProvider, unsortedMms)
	sort.Slice(mmsSortedByProvider, func(i, j int) bool {
		return mmsSortedByProvider[i].Provider < mmsSortedByProvider[j].Provider
	})
	return mmsSortedByProvider
}

func sortByCountry(unsortedMms []structs.MMSData) []structs.MMSData {
	sort.Slice(unsortedMms, func(i, j int) bool {
		return unsortedMms[i].Country[:2] < unsortedMms[j].Country[:2]
	})
	return unsortedMms
}
