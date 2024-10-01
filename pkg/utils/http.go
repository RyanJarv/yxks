package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func H[T any](w http.ResponseWriter, req *http.Request, f func(T) (any, error)) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("error reading request body: %v", err)
		return
	}

	t := *new(T)
	if err := json.Unmarshal(body, &t); err != nil {
		log.Printf("error unmarshalling request body: %v", err)
		return
	}

	response, err := f(t)
	if err != nil {
		log.Printf("error encrypting %T: %v", t, err)
		return
	}

	marshal, err := json.Marshal(response)
	if err != nil {
		log.Printf("marshalling response %T: %v", t, err)
		return
	}

	if _, err := w.Write(marshal); err != nil {
		log.Printf("error writing response %T: %v", t, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func TryJson(in any) string {
	marshal, err := json.Marshal(in)
	if err != nil {
		log.Printf("try json: %v", err)
		return ""
	}

	return string(marshal)
}

func TryJsonIndent(in any) string {
	marshal, err := json.MarshalIndent(in, "", "    ")
	if err != nil {
		log.Printf("try json indent: %v", err)
		return ""
	}

	return string(marshal)
}
