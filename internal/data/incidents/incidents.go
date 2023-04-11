package incidents

import (
	"diploma/internal/structs"
	"diploma/internal/support_functoins"
	"encoding/json"
	"io"
	"net/http"
	"sort"
)

func Incidents() ([]structs.IncidentData, string) {
	resp, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		return nil, support_functoins.ErrorToString(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, support_functoins.ErrorToString(err)
	}

	var incidents []structs.IncidentData
	if err := json.Unmarshal(body, &incidents); err != nil {
		return nil, support_functoins.ErrorToString(err)
	}

	sort.SliceStable(incidents, func(i, j int) bool {
		return incidents[i].Status < incidents[j].Status
	})

	return incidents, support_functoins.ErrorToString(err)
}
