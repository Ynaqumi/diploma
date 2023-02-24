package email

import (
	"io/ioutil"
	"log"
	"main/config"
	"main/status/check"
	"main/structs"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetEmail(emailChan chan map[string][][]structs.EmailData) {
	emailDataFile, err := os.Open(config.EmailDataFile)
	if err != nil {
		log.Println("Не удалось открыть файл email.data", err)
		emailChan <- nil
		return
	}
	defer emailDataFile.Close()

	readEmailDataFile, err := ioutil.ReadAll(emailDataFile)
	if err != nil {
		log.Println("Не удалось прочитать файл email.data", err)
	}

	emailCsvFile, err := os.Create(config.EmailCsvFile)
	if err != nil {
		log.Println("Не удалось создать файл email.csv", err)
	}

	var email []structs.EmailData
	line := strings.Split(string(readEmailDataFile), "\n")

	for i := 0; i < len(line); i++ {
		lineSlice := strings.Split(line[i], ";")
		deliveryTime, err := strconv.Atoi(lineSlice[2])
		if err != nil {
			continue
		}
		if len(lineSlice) == 3 && check.CountryCheck(lineSlice[0]) && check.ProviderEmail(lineSlice[1]) {
			emailCsvFile.WriteString(line[i])
			correctLine := structs.EmailData{Country: lineSlice[0], Provider: lineSlice[1], DeliveryTime: deliveryTime}
			email = append(email, correctLine)
		}
	}

	countriesMap := make(map[string][]structs.EmailData)
	countriesMap = createCountriesMap(countriesMap, email)
	for key := range countriesMap {
		countriesMap[key] = filterProvider(countriesMap[key])
	}

	newMap := make(map[string][][]structs.EmailData)
	for key := range countriesMap {
		fastestProviders := countriesMap[key][:3]
		slowestProviders := countriesMap[key][len(countriesMap[key])-4 : len(countriesMap[key])-1]
		var S [][]structs.EmailData
		S = append(S, fastestProviders)
		S = append(S, slowestProviders)
		newMap[key] = S
	}

	emailChan <- newMap
}

func filterProvider(emailSlice []structs.EmailData) []structs.EmailData {
	sort.SliceStable(emailSlice, func(i, j int) bool {
		return emailSlice[i].DeliveryTime > emailSlice[j].DeliveryTime
	})
	return emailSlice
}

func createCountriesMap(mapCountries map[string][]structs.EmailData, Email []structs.EmailData) map[string][]structs.EmailData {
	for i := 0; i < len(Email); i++ {
		circleCountry := Email[i].Country
		mapCountries[circleCountry] = append(mapCountries[circleCountry], Email[i])
	}
	return mapCountries
}
