package incidents

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

func Incidents() (incidents []IncidentData) {
	request, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		log.Printf("Не удалось выполнить GET запрос по MMS. Код ответа %v. Ошибка %v \n", request.StatusCode, err)
	} else {
		fmt.Printf("GET запрос по Incidents выполнен. Код ответа %v \n", request.StatusCode)
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Не удалось прочитать Get-запрос. Ошибка: %v", err)
	}

	if err := json.Unmarshal(body, &incidents); err != nil {
		log.Printf("Ошибка unmarshal %v", err)
	}
	return
}
