// Implements a kms xks server that implements the following APIs:
//
//	GetKeyMetadataHandler: /kms/xks/v1/keys/<externalKeyId>/metadata
//	Encrypt: /kms/xks/v1/keys/<externalKeyId>/encrypt
//	DecryptHandler: /kms/xks/v1/keys/<externalKeyId>/decrypt
//	GetHealthStatus: /kms/xks/v1/health
package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ryanjarv/yxks/pkg/handlers"
	"github.com/ryanjarv/yxks/pkg/middleware"
	"github.com/ryanjarv/yxks/pkg/utils"
	"github.com/samber/lo"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	debug    = flag.Bool("debug", false, "Enable debug logging")
	prefix   = flag.String("prefix", "yxks", "Prefix for handlers")
	listen   = flag.String("listen", "127.0.0.1:8443", "host to listen on")
	hostname = flag.String("hostname", "fno2i3na83.ryanjarv.sh", "hostname to use for letsencrypt")
	test     = flag.Bool("test", false, "testing mode")
	parse    = flag.Bool("parse", false, "parse the los")
)

func main() {
	ctx := utils.NewContext(context.Background())

	flag.Parse()

	if *debug {
		ctx = ctx.SetLoggingLevel(utils.DebugLogLevel)
	}

	if *parse {
		err := ParseLogs(ctx)
		if err != nil {
			log.Fatalf("error parsing logs: %v", err)
		}

	} else {
		log.Fatal(RunServer(ctx, *listen, *hostname, *test))
	}
}

func ParseLogs(ctx utils.Context) error {
	p := filepath.Join(lo.Must(os.UserHomeDir()), ".yxks", "logs")
	// list directory
	files, err := os.ReadDir(p)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		// read file
		fp := filepath.Join(p, f.Name())
		file, err := os.ReadFile(fp)
		if err != nil {
			return fmt.Errorf("error reading file: %v", err)
		}

		data := middleware.LogRequest{}
		if err = json.Unmarshal(file, &data); err != nil {
			return fmt.Errorf("error unmarshalling file: %v", err)
		}

		if data.Body != nil && data.Body["plaintext"] != nil {
			fmt.Println(data.Body["plaintext"])
			plaintext := data.Body["plaintext"].(string)
			derData, err := base64.StdEncoding.DecodeString(plaintext)
			if err != nil {
				return fmt.Errorf("error decoding plaintext: %v", err)
			}

			handlers.ParsePlaintextData(derData)
		}

	}
	return nil
}

func RunServer(ctx utils.Context, listen string, hostname string, testing bool) error {
	router := http.NewServeMux()

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

	if testing {
		ctx.Info.Printf("Starting testing server on %s", listen)
		if err := http.ListenAndServe(listen, router); err != nil {
			return fmt.Errorf("error starting server: %v", err)
		}
	} else {
		myRouter := middleware.SuperSecureMiddleware(ctx, router)
		myRouter = middleware.LoggingMiddleware(ctx, myRouter, *debug)

		cacheDir, err := os.UserCacheDir()
		if err != nil {
			return fmt.Errorf("error getting user cache dir: %v", err)
		}

		// Create a new autocert manager
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(hostname), // Replace with your domain
			Cache:      autocert.DirCache(filepath.Join(cacheDir, "yxks", "certs")),
		}
		server := &http.Server{
			Addr:      listen,
			Handler:   myRouter,
			TLSConfig: m.TLSConfig(),
		}

		ctx.Info.Printf("Starting server on %s", listen)
		if err := server.ListenAndServeTLS("", ""); err != nil {
			return fmt.Errorf("error starting server: %v", err)
		}
	}

	return nil
}
