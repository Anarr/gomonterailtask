package utils

import (
	"encoding/json"
	"net/http"
)

//Success return success response
func Success(w http.ResponseWriter, v interface{})  {
	res := make(map[string]interface{})
	res["data"] = v
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

//Error return error response
func Error(w http.ResponseWriter, code int, message string) {
	res := make(map[string]interface{})

	err := make(map[string]interface{})
	err["code"] = code
	err["message"] = message
	res["error"] = err

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}
