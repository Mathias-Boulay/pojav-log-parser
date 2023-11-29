// Package helloworld provides a set of Cloud Functions samples.
package parser

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("Entrypoint", Entrypoint)
}

// Entrypoint is an HTTP Cloud Function with a request parameter.
func Entrypoint(w http.ResponseWriter, r *http.Request) {
	// Limit the body size
	r.Body = http.MaxBytesReader(w, r.Body, 1000<<20) // MB to Bytes

	// Get the body then
	bodyText, _ := io.ReadAll(r.Body)

	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(ParseLog(string(bodyText), r.Header.Get("localize") == "true"))
	w.Write(json)
}
