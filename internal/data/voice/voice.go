package voice

import (
	"diploma/internal/structs"
	"diploma/internal/support_functoins"
	"os"
	"strconv"
	"strings"
)

func VoiceCall() (voiceCall []structs.VoiceCallData, error string) {
	data, err := os.ReadFile("simulator/voice.data")
	if err != nil {
		return voiceCall, support_functoins.ErrorToString(err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		lineStr := strings.Split(line, ";")
		if strings.Count(line, ";") == 7 && len(lineStr) == 8 && support_functoins.CountryCheck(lineStr[0]) && support_functoins.ProviderVoiceCheck(lineStr[3]) && support_functoins.BandwidthCheck(lineStr[1]) {
			stability64, _ := strconv.ParseFloat(lineStr[4], 32)
			stability32 := float32(stability64)
			TTFB, _ := strconv.Atoi(lineStr[5])
			voicePurity, _ := strconv.Atoi(lineStr[6])
			medianOfCallsTime, _ := strconv.Atoi(lineStr[7])
			voiceCall = append(voiceCall, structs.VoiceCallData{Country: lineStr[0], Bandwidth: lineStr[1], ResponseTime: lineStr[2], Provider: lineStr[3], ConnectionStability: stability32, TTFB: TTFB, VoicePurity: voicePurity, MedianOfCallsTime: medianOfCallsTime})
		}
	}
	return voiceCall, support_functoins.ErrorToString(err)
}
