package utils

import "net/http"

func GetExternalKeyId(req *http.Request) string {
	return req.PathValue("externalKeyId")
}
