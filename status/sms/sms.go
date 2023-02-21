package sms

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

func GetSms(smsChan chan [][]structs.SMSData) {
	smsDataFile, err := os.Open(config.SmsDataFile)
	if err != nil {
		log.Println("Не удалось открыть файл sms.data", err)
		smsChan <- nil
		return
	}
	defer smsDataFile.Close()

	readSmsDataFile, err := ioutil.ReadAll(smsDataFile)
	if err != nil {
		log.Println("Не удалось прочитать файл sms.data", err)
	}

	smsCsvFile, err := os.Create(config.SmsCsvFile)
	if err != nil {
		log.Println("Не удалось создать файл sms.csv", err)
	}

	var unsortedSms []structs.SMSData
	var sortedSms [][]structs.SMSData
	line := strings.Split(string(readSmsDataFile), "\n")

	for i := 0; i < len(line); i++ {
		lineSlice := strings.Split(line[i], ";")
		bandwidth, _ := strconv.Atoi(lineSlice[1])
		if len(lineSlice) == 4 && check.CountryCheck(lineSlice[0]) && check.ProviderSmsAndMms(lineSlice[3]) && (bandwidth >= 0 && bandwidth <= 100) {
			smsCsvFile.WriteString(line[i])
			correctLine := structs.SMSData{Country: lineSlice[0], Bandwidth: lineSlice[1], ResponseTime: lineSlice[1], Provider: lineSlice[3]}
			unsortedSms = append(unsortedSms, correctLine)
		}
	}

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
