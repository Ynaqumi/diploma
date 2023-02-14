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

func GetSms(smsChan chan [][]structs.SMSData) {
	DataFile, err := os.Open("C:/Users/touca/VSCode/diploma/src/simulator/sms.data")
	if err != nil {
		log.Println("Не удалось открыть файл", err)
		smsChan <- nil
		return
	}
	defer DataFile.Close()

	var unsortedSms []structs.SMSData

	scanner := bufio.NewScanner(DataFile)
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, ";")
		bandwidth, _ := strconv.Atoi(lineSlice[1])

		if len(lineSlice) == 4 &&
			lineSlice[2] != "" &&
			check.Country(lineSlice[0]) &&
			check.Provider(lineSlice[3]) &&
			(bandwidth >= 0 && bandwidth <= 100) {

			correctLine := structs.SMSData{lineSlice[0], lineSlice[1], lineSlice[2], lineSlice[3]}
			unsortedSms = append(unsortedSms, correctLine)
		}
	}
	var sortedSms [][]structs.SMSData

	smsSortedProvider := sortProvider(unsortedSms)
	sortedSms = append(sortedSms, smsSortedProvider)
	smsSortedCountry := sortCountry(unsortedSms)
	sortedSms = append(sortedSms, smsSortedCountry)
	smsChan <- sortedSms

	FileCSV, err := os.Create("C:/Users/touca/VSCode/diploma/pkg/data/sms.csv")
	if err != nil {
		log.Println("Не удалось создать файл", err)
	}
	defer FileCSV.Close()
}

func sortProvider(unsortedSms []structs.SMSData) []structs.SMSData {
	smsSortedProvider := make([]structs.SMSData, len(unsortedSms))
	copy(smsSortedProvider, unsortedSms)
	sort.Slice(smsSortedProvider, func(i, j int) bool {
		return smsSortedProvider[i].Provider < smsSortedProvider[j].Provider
	})
	return smsSortedProvider
}

func sortCountry(unsortedSms []structs.SMSData) []structs.SMSData {
	smsSortedCountry := make([]structs.SMSData, len(unsortedSms))
	copy(smsSortedCountry, unsortedSms)
	sort.Slice(smsSortedCountry, func(i, j int) bool {
		return smsSortedCountry[i].Country[:2] < smsSortedCountry[j].Country[:2]
	})
	return smsSortedCountry
}
