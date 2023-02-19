package voiceCall

import (
	"bufio"
	"log"
	"main/config"
	"main/status/check"
	"main/structs"
	"os"
	"strconv"
	"strings"
)

func GetVoice(voiceCallChan chan []structs.VoiceCallData) {
	voiceData, err := os.Open(config.VoiceCallDataFile)
	if err != nil {
		log.Println("Не удалось открыть файл Data", err)
		voiceCallChan <- nil
		return
	}
	defer voiceData.Close()

	VoiceCsvFile, err := os.Create(config.VoiceCallCsvFile)
	if err != nil {
		log.Println("Не удалось создать файл Csv", err)
	}
	defer VoiceCsvFile.Close()

	var voiceCall []structs.VoiceCallData

	scanner := bufio.NewScanner(voiceData)
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, ";")
		bandwidth, _ := strconv.Atoi(lineSlice[1])
		Stability64, _ := strconv.ParseFloat(lineSlice[3], 32)
		Stability32 := float32(Stability64)
		StabilityInt := int(Stability32)
		TTFB, _ := strconv.Atoi(lineSlice[4])
		VoicePurity, _ := strconv.Atoi(lineSlice[5])
		MedianOfCallsTime, _ := strconv.Atoi(lineSlice[6])

		if len(lineSlice) == 8 && lineSlice[2] != "" && check.CountrySmsAndMms(lineSlice[0]) && check.ProviderSmsAndMms(lineSlice[3]) && (bandwidth >= 0 && bandwidth <= 100) && Stability32 != 0 && TTFB != 0 && VoicePurity != 0 && MedianOfCallsTime != 0 {
			VoiceCsvFile.WriteString(lineSlice[0] + ";" + lineSlice[1] + ";" + lineSlice[2] + ";" + lineSlice[3] + string(StabilityInt) + lineSlice[4] + lineSlice[5] + lineSlice[6] + "\n")
			correctLine := structs.VoiceCallData{Country: lineSlice[0], Bandwidth: lineSlice[1], ResponseTime: lineSlice[2], Provider: lineSlice[3], ConnectionStability: Stability32, TTFB: TTFB, VoicePurity: VoicePurity, MedianOfCallsTime: MedianOfCallsTime}
			voiceCall = append(voiceCall, correctLine)
		}
	}
	voiceCallChan <- voiceCall
}
