package voice

import (
	"diploma2/status/check"
	"log"
	"os"
	"strconv"
	"strings"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

func VoiceCall() (voiceCall []VoiceCallData) {
	data, err := os.ReadFile("simulator/voice.data")
	if err != nil {
		log.Printf("Не удалось прочитать файл voice.data. Ошибка: %v", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		lineStr := strings.Split(line, ";")
		if strings.Count(line, ";") == 7 && len(lineStr) == 8 && check.CountryCheck(lineStr[0]) && check.ProviderVoiceCheck(lineStr[3]) && check.BandwidthCheck(lineStr[1]) {
			stability64, _ := strconv.ParseFloat(lineStr[4], 32)
			stability32 := float32(stability64)
			TTFB, _ := strconv.Atoi(lineStr[5])
			voicePurity, _ := strconv.Atoi(lineStr[6])
			medianOfCallsTime, _ := strconv.Atoi(lineStr[7])
			voiceCall = append(voiceCall, VoiceCallData{Country: lineStr[0], Bandwidth: lineStr[1], ResponseTime: lineStr[2], Provider: lineStr[3], ConnectionStability: stability32, TTFB: TTFB, VoicePurity: voicePurity, MedianOfCallsTime: medianOfCallsTime})
		}
	}
	return
}
