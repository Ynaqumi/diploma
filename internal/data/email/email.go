package email

import (
	"diploma/internal/check"
	"diploma/internal/structs"
	"log"
	"os"
	"strconv"
	"strings"
)

func Email() (email []structs.EmailData) {
	dataFileContent, err := os.ReadFile("simulator/email.data")
	if err != nil {
		log.Printf("Не удалось прочитать файл email.data. Ошибка: %v", err)
	}

	for _, line := range strings.Split(string(dataFileContent), "\n") {
		if strings.Count(line, ";") == 2 {
			lineStr := strings.Split(line, ";")
			deliveryTime, err := strconv.Atoi(lineStr[2])
			if err != nil || !check.CountryCheck(lineStr[0]) || !check.ProviderEmail(lineStr[1]) {
				continue
			}
			email = append(email, structs.EmailData{Country: lineStr[0] + ";", Provider: lineStr[1] + ";", DeliveryTime: deliveryTime})
		}
	}
	return
}
