package voiceCall

import (
	"io/ioutil"
	"log"
	"main/config"
	"main/status/check"
	"main/structs"
	"os"
	"strconv"
	"strings"
)

func GetVoice(voiceCallChan chan []structs.VoiceCallData) {
	voiceDataFile, err := os.Open(config.VoiceCallDataFile)
	if err != nil {
		log.Println("Не удалось открыть файл voice.data", err)
		voiceCallChan <- nil
		return
	}
	defer voiceDataFile.Close()

	readVoiceDataFile, err := ioutil.ReadAll(voiceDataFile)
	if err != nil {
		log.Println("Не удалось прочитать файл voice.data", err)
	}

	voiceCsvFile, err := os.Create(config.VoiceCallCsvFile)
	if err != nil {
		log.Println("Не удалось создать файл voicecall.csv", err)
	}

	var voiceCall []structs.VoiceCallData
	line := strings.Split(string(readVoiceDataFile), "\n")

	for i := 0; i < len(line); i++ {
		lineSlice := strings.Split(line[i], ";")
		bandwidth, err := strconv.Atoi(lineSlice[1])
		if err != nil {
			continue
		}
		Stability64, err := strconv.ParseFloat(lineSlice[4], 32)
		if err != nil {
			continue
		}
		Stability32 := float32(Stability64)
		TTFB, err := strconv.Atoi(lineSlice[5])
		if err != nil {
			continue
		}
		VoicePurity, err := strconv.Atoi(lineSlice[6])
		if err != nil {
			continue
		}
		MedianOfCallsTime, err := strconv.Atoi(lineSlice[7])
		if err != nil {
			continue
		}

		if len(lineSlice) == 8 && check.CountryCheck(lineSlice[0]) && check.ProviderVoiceCall(lineSlice[3]) && (bandwidth >= 0 && bandwidth <= 100) && Stability32 != 0 && TTFB != 0 && VoicePurity != 0 && MedianOfCallsTime != 0 {
			voiceCsvFile.WriteString(line[i])
			correctLine := structs.VoiceCallData{Country: lineSlice[0], Bandwidth: lineSlice[1], ResponseTime: lineSlice[2], Provider: lineSlice[3], ConnectionStability: Stability32, TTFB: TTFB, VoicePurity: VoicePurity, MedianOfCallsTime: MedianOfCallsTime}
			voiceCall = append(voiceCall, correctLine)
		}
	}

	voiceCallChan <- voiceCall
}
