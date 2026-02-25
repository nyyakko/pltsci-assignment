package hoover

import (
	"assignment/api/hoover/contracts/requests"
	"assignment/utils"
	"encoding/json"
	"net/http"
)

type controller struct {}

var Controller controller

func (controller) CleaningSessions(w http.ResponseWriter, r *http.Request) *http_utils.HttpError {
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	w.Header().Set("Content-Type", "application/json")

	var request requests.CleaningRequest
	decoderErr := decoder.Decode(&request)
	if decoderErr != nil {
		return http_utils.MakeInternalServerError("JSON_DECODER_FAILED", decoderErr.Error())
	}

	res, serviceErr := Service.CleaningSessions(request)
	if serviceErr != nil {
		return serviceErr
	}

	encoder.Encode(res)

	return nil
}
