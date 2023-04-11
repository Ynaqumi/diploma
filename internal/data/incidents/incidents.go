package incidents

import (
	"diploma/internal/structs"
	"diploma/internal/support_functoins"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Incidents() (incidents []structs.IncidentData, error string) {
	request, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		log.Printf("Не удалось выполнить GET-запрос по Incidents. Код ответа %v. Ошибка %v \n", request.StatusCode, err)
		return incidents, support_functoins.ErrorToString(err)
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Не удалось прочитать GET-запрос по Incidents. Ошибка: %v", err)
		return incidents, support_functoins.ErrorToString(err)
	}

	if err := json.Unmarshal(body, &incidents); err != nil {
		log.Printf("Ошибка unmarshal по Incidents: %v", err)
		return incidents, support_functoins.ErrorToString(err)
	}
	return incidents, support_functoins.ErrorToString(err)
}
