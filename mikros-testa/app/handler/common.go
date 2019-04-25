package handler

import (
	"fmt"
	"net/http"
	"strings"
)

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func GetProjectName(r *http.Request) string {
	headers := make(map[string]interface{})
	for k, v := range r.Header {
		headers[strings.ToLower(k)] = string(v[0])
	}
	if value, ok := headers["project"]; ok {
		return value.(string)
	}
	return ""
}
