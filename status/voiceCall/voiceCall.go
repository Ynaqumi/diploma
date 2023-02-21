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
		bandwidth, _ := strconv.Atoi(lineSlice[1])
		Stability64, _ := strconv.ParseFloat(lineSlice[4], 32)
		Stability32 := float32(Stability64)
		TTFB, _ := strconv.Atoi(lineSlice[5])
		VoicePurity, _ := strconv.Atoi(lineSlice[6])
		MedianOfCallsTime, _ := strconv.Atoi(lineSlice[7])

		if len(lineSlice) == 8 && check.CountryCheck(lineSlice[0]) && check.ProviderVoiceCall(lineSlice[3]) && (bandwidth >= 0 && bandwidth <= 100) && Stability32 != 0 && TTFB != 0 && VoicePurity != 0 && MedianOfCallsTime != 0 {
			voiceCsvFile.WriteString(line[i])
			correctLine := structs.VoiceCallData{Country: lineSlice[0], Bandwidth: lineSlice[1], ResponseTime: lineSlice[2], Provider: lineSlice[3], ConnectionStability: Stability32, TTFB: TTFB, VoicePurity: VoicePurity, MedianOfCallsTime: MedianOfCallsTime}
			voiceCall = append(voiceCall, correctLine)
		}
	}

	voiceCallChan <- voiceCall
}
