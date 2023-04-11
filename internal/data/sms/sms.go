package sms

import (
	"diploma/internal/structs"
	"diploma/internal/support_functoins"
	"log"
	"os"
	"sort"
	"strings"
)

func Sms() ([][]structs.SMSData, string) {
	data, err := os.ReadFile("simulator/sms.data")
	if err != nil {
		log.Printf("Не удалось прочитать файл sms.data. Ошибка: %v", err)
		return nil, support_functoins.ErrorToString(err)
	}

	unsortedSms := []structs.SMSData{}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.Count(line, ";") == 3 {
			lineStr := strings.Split(line, ";")
			if support_functoins.CountryCheck(lineStr[0]) && support_functoins.BandwidthCheck(lineStr[1]) && support_functoins.ProviderSmsAndMMSCheck(lineStr[3]) {
				unsortedSms = append(unsortedSms, structs.SMSData{lineStr[0] + ";", lineStr[1] + ";", lineStr[2] + ";", lineStr[3]})
			}
		}
	}

	smsFullCountry := makeFullCountry(unsortedSms)
	smsSortedByProvider := sortByProvider(smsFullCountry)
	smsSortedByCountry := sortByCountry(smsFullCountry)

	return [][]structs.SMSData{smsSortedByProvider, smsSortedByCountry}, support_functoins.ErrorToString(err)
}

func makeFullCountry(shortnameCountry []structs.SMSData) []structs.SMSData {
	for i := range shortnameCountry {
		if fullCountry, ok := support_functoins.CountriesList[shortnameCountry[i].Сountry]; ok {
			shortnameCountry[i].Сountry = fullCountry
		}
	}
	return shortnameCountry
}

func sortByProvider(unsortedSmsCollection []structs.SMSData) []structs.SMSData {
	smsSortedByProvider := make([]structs.SMSData, len(unsortedSmsCollection))
	copy(smsSortedByProvider, unsortedSmsCollection)
	sort.Slice(smsSortedByProvider, func(i, j int) bool {
		return smsSortedByProvider[i].Provider < smsSortedByProvider[j].Provider
	})
	return smsSortedByProvider
}

func sortByCountry(smsFullCountry []structs.SMSData) []structs.SMSData {
	sort.Slice(smsFullCountry, func(i, j int) bool {
		return smsFullCountry[i].Сountry[:2] < smsFullCountry[j].Сountry[:2]
	})
	return smsFullCountry
}
