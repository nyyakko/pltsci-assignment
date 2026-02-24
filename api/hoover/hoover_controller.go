package hoover

import (
	"assignment/api/hoover/contracts/requests"
	"encoding/json"
	"net/http"
)

type controller struct {}

var Controller controller

func (controller) CleaningSessions(w http.ResponseWriter, r *http.Request) error {
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	var request requests.CleaningRequest
	err := decoder.Decode(&request)
	if err != nil {
		return err
	}

	res, err := Service.CleaningSessions(request)
	if err != nil {
		return err
	}

	encoder.Encode(res)

	return nil
}
