package sms

import (
	"bufio"
	"fmt"
	"log"
	"main/config"
	"main/status/check"
	"main/structs"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetSms(smsChan chan [][]structs.SMSData) {
	SmsDataFile, err := os.Open(config.SmsDataFile)
	if err != nil {
		log.Println("Не удалось открыть файл Data", err)
		smsChan <- nil
		return
	}
	defer SmsDataFile.Close()

	SmsCsvFile, err := os.Create(config.SmsCsvFile)
	if err != nil {
		log.Println("Не удалось создать файл Csv", err)
	}
	defer SmsCsvFile.Close()

	var unsortedSms []structs.SMSData

	scanner := bufio.NewScanner(SmsDataFile)
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, ";")
		bandwidth, _ := strconv.Atoi(lineSlice[1])

		if len(lineSlice) == 4 && lineSlice[2] != "" && check.CountrySmsAndMms(lineSlice[0]) && check.ProviderSmsAndMms(lineSlice[3]) && (bandwidth >= 0 && bandwidth <= 100) {
			SmsCsvFile.WriteString(lineSlice[0] + ";" + lineSlice[1] + ";" + lineSlice[2] + ";" + lineSlice[3] + "\n")
			if err != nil {
				fmt.Println(err.Error())
			}
			correctLine := structs.SMSData{Country: lineSlice[0], Bandwidth: lineSlice[1], ResponseTime: lineSlice[1], Provider: lineSlice[3]}
			unsortedSms = append(unsortedSms, correctLine)
		}
	}
	var sortedSms [][]structs.SMSData

	smsSortedProvider, smsSortedCountry := sortProviderSms(unsortedSms), sortCountrySms(unsortedSms)
	sortedSms = append(sortedSms, smsSortedProvider)
	sortedSms = append(sortedSms, smsSortedCountry)
	smsChan <- sortedSms

}

func sortProviderSms(unsortedSms []structs.SMSData) []structs.SMSData {
	smsSortedProvider := make([]structs.SMSData, len(unsortedSms))
	copy(smsSortedProvider, unsortedSms)
	sort.Slice(smsSortedProvider, func(i, j int) bool {
		return smsSortedProvider[i].Provider < smsSortedProvider[j].Provider
	})
	return smsSortedProvider
}

func sortCountrySms(unsortedSms []structs.SMSData) []structs.SMSData {
	smsSortedCountry := make([]structs.SMSData, len(unsortedSms))
	copy(smsSortedCountry, unsortedSms)
	sort.Slice(smsSortedCountry, func(i, j int) bool {
		return smsSortedCountry[i].Country[:2] < smsSortedCountry[j].Country[:2]
	})
	return smsSortedCountry
}
