package status

import (
	"bufio"
	"log"
	"main/pkg/check"
	"main/pkg/structs"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetEmail(emailChan chan map[string][][]structs.EmailData) {
	emailData, err := os.Open("C:/Users/touca/VSCode/diploma/src/simulator/email.data")
	if err != nil {
		log.Println("Не удалось открыть файл", err)
		emailChan <- nil
		return
	}
	defer emailData.Close()

	var email []structs.EmailData

	scanner := bufio.NewScanner(emailData)
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, ";")

		if len(lineSlice) == 3 &&
			lineSlice[2] != "" &&
			check.Country(lineSlice[0]) &&
			check.ProviderEmail(lineSlice[1]) {

			deliveryTime, _ := strconv.Atoi(lineSlice[2])
			correctLine := structs.EmailData{lineSlice[0], lineSlice[1], deliveryTime}
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
