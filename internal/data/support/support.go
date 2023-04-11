package support

import (
	"diploma/internal/structs"
	"diploma/internal/support_functoins"
	"encoding/json"
	"io"
	"net/http"
)

func Support() (support []int, error string) {
	request, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {
		return support, support_functoins.ErrorToString(err)
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return support, support_functoins.ErrorToString(err)
	}

	supportData := []structs.SupportData{}
	if err := json.Unmarshal(body, &supportData); err != nil {
		return support, support_functoins.ErrorToString(err)
	}

	load := 0
	for _, data := range supportData {
		load += data.ActiveTickets
	}

	var loadLvl int
	if load < 9 {
		loadLvl = 1
	} else if load > 16 {
		loadLvl = 3
	} else {
		loadLvl = 2
	}

	support = append(support, loadLvl)

	ans := (float64(load)/2 + 0.5) * 3.333333333333333
	support = append(support, int(ans+1))

	return support, support_functoins.ErrorToString(err)
}
