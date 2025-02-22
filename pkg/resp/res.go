package resp

import (
	"encoding/json"
	"net/http"
)

func ResponceJson(w http.ResponseWriter, respCode int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respCode)
	json.NewEncoder(w).Encode(resp)
}
