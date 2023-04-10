package sms

import (
	"diploma/internal/check"
	"diploma/internal/structs"
	"log"
	"os"
	"strings"
)

func Sms() (sms []structs.SMSData) {
	data, err := os.ReadFile("simulator/sms.data")
	if err != nil {
		log.Printf("Не удалось прочитать файл sms.data. Ошибка: %v", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.Count(line, ";") == 3 {
			lineStr := strings.Split(line, ";")
			if check.CountryCheck(lineStr[0]) && check.BandwidthCheck(lineStr[1]) && check.ProviderSmsAndMMSCheck(lineStr[3]) {
				sms = append(sms, structs.SMSData{lineStr[0] + ";", lineStr[1] + ";", lineStr[2] + ";", lineStr[3]})
			}
		}
	}
	return
}
