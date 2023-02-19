package incidents

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/structs"
	"net/http"
)

func GetIncidents(incidentsChan chan []structs.IncidentData) {
	request, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		log.Println("Не удалось выполнить GET запрос", err)
		incidentsChan <- nil
		return
	}

	if request.StatusCode != 200 {
		var IncidentDataCollection []structs.IncidentData
		log.Println("Ошибка статус код Incidents не равен 200", IncidentDataCollection)
		incidentsChan <- nil
		return
	} else {
		fmt.Println("Код ответа 200")
	}

	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println("Не удалось прочитать запрос", err)
		incidentsChan <- nil
		return
	}

	var incidentData []structs.IncidentData
	if err := json.Unmarshal(bytes, &incidentData); err != nil {
		log.Println("Ошибка unmarshal", err)
		incidentsChan <- nil
		return
	}

	incidentsChan <- incidentData
}
