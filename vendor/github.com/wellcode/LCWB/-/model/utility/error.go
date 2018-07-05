package utility

import (
	"encoding/json"
	"net/http"
)

//error: "qwertyuiop"
//description : "asdfghjkl;"
//fields :
// field : "dpp"
// description : "qweqwe"

func ErrorResonseJSON(w http.ResponseWriter, r *http.Request, errorDescription string, array [](map[string]string)) {
	result := make(map[string]interface{})
	result["error"] = true
	result["description"] = errorDescription
	result["fields"] = array
	jsonData, _ := json.Marshal(result)
	w.WriteHeader(500)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	return
}
