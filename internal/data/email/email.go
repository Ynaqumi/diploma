package email

import (
	"diploma/internal/structs"
	"diploma/internal/support_functoins"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Email() (map[string][][]structs.EmailData, string) {
	dataFileContent, err := os.ReadFile("simulator/email.data")
	if err != nil {
		log.Printf("Не удалось прочитать файл email.data. Ошибка: %v", err)
		return nil, support_functoins.ErrorToString(err)
	}

	email := []structs.EmailData{}
	for _, line := range strings.Split(string(dataFileContent), "\n") {
		if strings.Count(line, ";") == 2 {
			lineStr := strings.Split(line, ";")
			deliveryTime, err := strconv.Atoi(lineStr[2])
			if err != nil || !support_functoins.CountryCheck(lineStr[0]) || !support_functoins.ProviderEmail(lineStr[1]) {
				continue
			}
			email = append(email, structs.EmailData{Country: lineStr[0], Provider: lineStr[1], DeliveryTime: deliveryTime})
		}
	}

	countriesMap := make(map[string][]structs.EmailData)
	for _, data := range email {
		countriesMap[data.Country] = append(countriesMap[data.Country], data)
	}

	sortedEmail := make(map[string][][]structs.EmailData)
	for key := range countriesMap {
		sort.SliceStable(countriesMap[key], func(i, j int) bool {
			return countriesMap[key][i].DeliveryTime > countriesMap[key][j].DeliveryTime
		})
		fastestProviders := countriesMap[key][:3]
		slovestProviders := countriesMap[key][len(countriesMap[key])-4 : len(countriesMap[key])-1]
		cercleSloce := [][]structs.EmailData{}
		cercleSloce = append(cercleSloce, fastestProviders)
		cercleSloce = append(cercleSloce, slovestProviders)
		sortedEmail[key] = cercleSloce
	}
	return sortedEmail, support_functoins.ErrorToString(err)
}
