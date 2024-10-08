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
	"strings"
)

var (
	debug  = flag.Bool("debug", false, "Enable debug logging")
	prefix = flag.String("prefix", "yxks", "Prefix for handlers")
)

func main() {
	ctx := utils.NewContext(context.Background())

	if *debug {
		ctx = ctx.SetLoggingLevel(utils.DebugLogLevel)
	}

	log.Fatal(RunServer(ctx))
}

func RunServer(ctx utils.Context) error {
	router := http.NewServeMux()
	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if strings.HasPrefix(*prefix, "/") {
		*prefix = "yxks"
	} else {
		*prefix = "/" + *prefix
	}

	router.HandleFunc(*prefix+"/kms/xks/v1/health", handlers.HealthHandler)
	router.HandleFunc(*prefix+"/kms/xks/v1/keys/{externalKeyId}/metadata", handlers.GetKeyMetadataHandler)
	router.HandleFunc(*prefix+"/kms/xks/v1/keys/{externalKeyId}/encrypt", handlers.EncryptHandler)
	router.HandleFunc(*prefix+"/kms/xks/v1/keys/{externalKeyId}/decrypt", handlers.DecryptHandler)
	router.HandleFunc("/", handlers.GetDefaultHandler(ctx))

	return s.ListenAndServe()
}
