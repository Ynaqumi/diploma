package status

import (
	"bufio"
	"log"
	"main/pkg/check"
	"main/pkg/structs"
	"os"
	"strconv"
	"strings"
)

const VoiceCallCsv = "diploma/pkg/data/voicecall.csv"

func GetVoice(voiceCallChan chan []structs.VoiceCallData) {
	voiceData, err := os.Open(VoiceCallCsv)
	if err != nil {
		log.Println("Не удалось открыть файл", err)
		voiceCallChan <- nil
		return
	}
	defer voiceData.Close()

	var voiceCall []structs.VoiceCallData

	scanner := bufio.NewScanner(voiceData)
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, ";")
		bandwidth, _ := strconv.Atoi(lineSlice[1])
		Stability, _ := strconv.ParseFloat(lineSlice[4], 32)
		Stability32 := float32(Stability)
		TTFB, _ := strconv.Atoi(lineSlice[5])
		VoicePurity, _ := strconv.Atoi(lineSlice[6])
		MedianOfCallsTime, _ := strconv.Atoi(lineSlice[7])

		if len(lineSlice) == 8 &&
			lineSlice[2] != "" &&
			check.Country(lineSlice[0]) &&
			check.Provider(lineSlice[3]) &&
			(bandwidth >= 0 && bandwidth <= 100) &&
			Stability32 != 0 &&
			TTFB != 0 &&
			VoicePurity != 0 &&
			MedianOfCallsTime != 0 {
			correctLine := structs.VoiceCallData{lineSlice[0], lineSlice[1], lineSlice[2], lineSlice[3], Stability32, TTFB, VoicePurity, MedianOfCallsTime}
			voiceCall = append(voiceCall, correctLine)
		}
	}
	voiceCallChan <- voiceCall
}
