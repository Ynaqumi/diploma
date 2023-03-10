package support

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/structs"
	"net/http"
)

func GetSupport(supportChan chan []int) {
	respAns, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {
		log.Println("Не удалось выполнить GET запрос", err)
		supportChan <- nil
		return
	}

	if respAns.StatusCode != 200 {
		var supportDataCollection []structs.SupportData
		log.Println("Ошибка статус код Support не равен 200", supportDataCollection)
		supportChan <- nil
		return
	} else {
		fmt.Println("Код ответа 200")
	}

	byteAns, err := ioutil.ReadAll(respAns.Body)
	if err != nil {
		log.Println("Не удалось прочитать запрос", err)
		supportChan <- nil
		return
	}

	var supportData []structs.SupportData
	if err := json.Unmarshal(byteAns, &supportData); err != nil {
		log.Println("Ошибка unmarshal", err)
		supportChan <- nil
		return
	}

	load := 0
	for i := 0; i < len(supportData); i++ {
		load = load + supportData[i].ActiveTickets
	}
	log.Println(load)
	loadLvl := 0
	if load < 9 {
		loadLvl = 1
	} else if load > 16 {
		loadLvl = 3
	} else {
		loadLvl = 2
	}

	var supportStatus []int
	supportStatus = append(supportStatus, loadLvl)

	if load%2 == 1 {
		ans := (float64(load)/2 + 1) * 3.333333333333333
		if ans-float64(int(ans)) == 0 {
			supportStatus = append(supportStatus, int(ans))
		} else {
			supportStatus = append(supportStatus, int(ans+1))
		}
	} else {
		ans := (float64(load)/2 + 0.5) * 3.333333333333333
		if ans-float64(int(ans)) == 0 {
			supportStatus = append(supportStatus, int(ans))
		} else {
			supportStatus = append(supportStatus, int(ans+1))
		}
	}
	supportChan <- supportStatus
}
