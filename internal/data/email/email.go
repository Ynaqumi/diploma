package email

import (
	"diploma/internal/structs"
	"diploma/internal/support_functoins"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Email() (map[string][][]structs.EmailData, string) {
	dataFileContent, err := os.ReadFile("simulator/email.data")
	if err != nil {
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

	sortedEmail := make(map[string][][]structs.EmailData)
	for _, data := range email {
		sort.SliceStable(email, func(i, j int) bool {
			return email[i].DeliveryTime > email[j].DeliveryTime
		})
		fastestProviders := email[:3]
		slovestProviders := email[len(email)-4 : len(email)-1]
		cercleSloce := [][]structs.EmailData{}
		cercleSloce = append(cercleSloce, fastestProviders)
		cercleSloce = append(cercleSloce, slovestProviders)
		sortedEmail[data.Country] = cercleSloce
	}
	return sortedEmail, support_functoins.ErrorToString(err)
}
