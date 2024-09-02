package handlers

import (
	"fmt"
	"github.com/ryanjarv/yxks/pkg/utils"
	"github.com/samber/lo"
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

		ctx.Info.Printf("default handler: %s %s %s", req.Method, req.URL.Path, body)
		lo.Must(fmt.Fprintf(w, "Called default handler with: %s", body))
	}
}
