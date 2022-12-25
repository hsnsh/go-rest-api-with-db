package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/HsnCorp/go-hsn-library/logger"
	"net/http"
	"strings"
)

type ResponseDto struct {
	Success  bool        `json:"success"`
	Messages []string    `json:"messages"`
	Data     interface{} `json:"data"`
}

func RespondErrorWithMessage(w http.ResponseWriter, message string) {
	RespondErrorWithMessages(w, strings.Split(message, "\n"))
}

func RespondErrorWithMessages(w http.ResponseWriter, messages []string) {
	respondBaseWithData(w, http.StatusBadRequest, ResponseDto{Success: false, Messages: messages})
}

func RespondOkWithData(w http.ResponseWriter, payload interface{}) {
	respondBaseWithData(w, http.StatusOK, ResponseDto{Success: true, Data: payload})
}

func RespondCreatedWithData(w http.ResponseWriter, payload interface{}) {
	respondBaseWithData(w, http.StatusCreated, ResponseDto{Success: true, Data: payload})
}

func RespondOk(w http.ResponseWriter) {
	RespondWithCode(w, http.StatusOK)
}

func RespondWithCode(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

func respondBaseWithData(w http.ResponseWriter, code int, apiResponse ResponseDto) {

	newPayload := NilSliceToEmptySlice(apiResponse)
	response, err := json.Marshal(newPayload)
	if err != nil {
		RespondWithCode(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func HandlePanicAndRecovery(w http.ResponseWriter, appLogger logger.IFileLogger) {
	if r := recover(); r != nil {
		appLogger.Warning(fmt.Sprintf("%s", r))
		RespondWithCode(w, http.StatusInternalServerError)
	}
}
