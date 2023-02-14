package check

import (
	"main/pkg/check/checkList"
)

func Country(str string) bool {
	for i := 0; i < len(checkList.Countries); i++ {
		if str == checkList.Countries[string(i)] {
			return true
		}
	}
	return false
}

func Provider(str string) bool {
	for i := 0; i < len(checkList.ProvidersListSMSAndMMS); i++ {
		if str == checkList.ProvidersListSMSAndMMS[i] {
			return true
		}
	}
	return false
}

func ProviderEmail(str string) bool {
	for i := 0; i < len(checkList.ProvidersListEmail); i++ {
		if str == checkList.ProvidersListEmail[i] {
			return true
		}
	}
	return false
}
