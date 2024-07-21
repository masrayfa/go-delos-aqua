package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(req *http.Request, result interface{}) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(result)
	if err != nil {
		return err
	}
	
	return nil
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}