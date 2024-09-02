// Implements a kms xks server that implements the following APIs:
//
//	GetKeyMetadataHandler: /kms/xks/v1/keys/<externalKeyId>/metadata
//	Encrypt: /kms/xks/v1/keys/<externalKeyId>/encrypt
//	DecryptHandler: /kms/xks/v1/keys/<externalKeyId>/decrypt
//	GetHealthStatus: /kms/xks/v1/health
package main

import (
	"context"
	"flag"
	"github.com/ryanjarv/yxks/pkg/handlers"
	"github.com/ryanjarv/yxks/pkg/utils"
	"log"
	"net/http"
)

var (
	debug = flag.Bool("debug", false, "Enable debug logging")
)

func main() {
	ctx := utils.NewContext(context.Background())

	if *debug {
		ctx = ctx.SetLoggingLevel(utils.DebugLogLevel)
	}

	log.Fatal(RunServer(ctx))
}

func RunServer(ctx utils.Context) error {
	http.HandleFunc("/kms/xks/v1/health", handlers.HealthHandler)
	http.HandleFunc("/kms/xks/v1/keys/{externalKeyId}/metadata", handlers.GetKeyMetadataHandler)
	http.HandleFunc("/kms/xks/v1/keys/{externalKeyId}/encrypt", handlers.EncryptHandler)
	http.HandleFunc("/kms/xks/v1/keys/{externalKeyId}/decrypt", handlers.DecryptHandler)

	return http.ListenAndServe("localhost:8080", nil)
}
