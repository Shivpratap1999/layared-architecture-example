package response

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, code int, v interface{}) {
	bytes, err := json.Marshal(v)
	if err != nil {
		SendError(w, http.StatusInternalServerError, "Somthing Wents Wrong !")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(bytes)

}
func SendError(w http.ResponseWriter, code int, msg string) {
	SendJson(w, code, map[string]string{"error": msg})
}
