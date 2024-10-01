package handlers

import (
	"github.com/ryanjarv/yxks/pkg/utils"
	"io"
	"net/http"
)

// GetDefaultHandler endpoint
// URI: /
func GetDefaultHandler(ctx utils.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		ctx.Error.Printf("default handler called: %s %s %s", req.Method, req.URL.Path, body)

		http.NotFound(w, req)
	}
}
