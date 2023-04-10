package email

import (
	"diploma2/status/check"
	"log"
	"os"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func Email() (email []EmailData) {
	dataFileContent, err := os.ReadFile("simulator/email.data")
	if err != nil {
		log.Println("Не удалось прочитать email.data файл", err)
	}

	for _, line := range strings.Split(string(dataFileContent), "\n") {
		if strings.Count(line, ";") == 2 {
			lineStr := strings.Split(line, ";")
			deliveryTime, err := strconv.Atoi(lineStr[2])
			if err != nil || !check.CountryCheck(lineStr[0]) || !check.ProviderEmail(lineStr[1]) {
				continue
			}
			email = append(email, EmailData{Country: lineStr[0] + ";", Provider: lineStr[1] + ";", DeliveryTime: deliveryTime})
		}
	}
	return
}
