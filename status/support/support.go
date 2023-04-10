package support

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func Support() (support []SupportData) {
	request, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {
		log.Printf("Не удалось выполнить GET запрос по MMS. Код ответа %v. Ошибка %v \n", request.StatusCode, err)
	} else {
		fmt.Printf("GET запрос по Support выполнен. Код ответа %v \n", request.StatusCode)
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println("Не удалось прочитать get-запрос", err)
	}

	if err := json.Unmarshal(body, &support); err != nil {
		log.Println("Ошибка unmarshal", err)
	}
	return
}
