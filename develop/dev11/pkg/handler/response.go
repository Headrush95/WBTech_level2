package handler

import (
	"dev11/models"
	"encoding/json"
	"net/http"
)

type ErrorResp struct {
	Error string `json:"error"`
}

type SuccessGetResp struct {
	Result []models.Event `json:"result"`
}

type SuccessPostDeleteResp struct {
	Result string `json:"result"`
}

func SuccessGetResponse(w http.ResponseWriter, events []models.Event) {
	output := SuccessGetResp{
		Result: events,
	}

	resp, err := json.MarshalIndent(output, " ", "")
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
}

func SuccessPostDeleteResponse(w http.ResponseWriter, message string) {
	resp, _ := json.MarshalIndent(SuccessPostDeleteResp{Result: message}, " ", "")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}

func ErrorResponse(w http.ResponseWriter, err string, statusCode int) {
	resp, _ := json.MarshalIndent(ErrorResp{Error: err}, " ", "")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(resp)
}
