package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ryanjarv/yxks/pkg/utils"
	"github.com/samber/lo"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// LoggingMiddleware logs the details of each incoming request
func LoggingMiddleware(ctx utils.Context, next http.Handler, debug bool) http.Handler {
	logDir := filepath.Join(lo.Must(os.UserHomeDir()), ".yxks", "logs")
	lo.Must0(os.MkdirAll(logDir, os.ModePerm))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if debug {
			// Log request details
			log.Printf("[DEBUG] received request: %s %s from %s", r.Method, r.URL, r.RemoteAddr)

			request, err := httputil.DumpRequest(r, true)
			if err != nil {
				ctx.Error.Println("Error dumping request: ", err)
			}
			fmt.Println("\n\n" + string(request) + "\n\n")
		}

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
		}
		r.Body = io.NopCloser(bytes.NewReader(reqBody))

		// Create a response writer to capture response details
		rw := NewResponseLogger(w)

		// Call the next handler
		next.ServeHTTP(rw, r)

		if err := LogBody(r, reqBody, rw.body, logDir); err != nil {
			log.Printf("Error logging request body: %v", err)
		}
	})
}

type LogRequest struct {
	Host     string         `json:"host"`
	Method   string         `json:"method"`
	URL      string         `json:"url"`
	Headers  http.Header    `json:"headers"`
	Body     map[string]any `json:"body"`
	Response map[string]any `json:"response"`
}

func LogBody(r *http.Request, reqBodyData []byte, respBodyData []byte, logDir string) error {
	reqBody := map[string]any{}
	if err := json.Unmarshal(reqBodyData, &reqBody); err != nil {
		return fmt.Errorf("error marshalling request: %v, value: %s", err, reqBodyData)
	}

	respBody := map[string]any{}
	if err := json.Unmarshal(respBodyData, &respBody); err != nil {
		return fmt.Errorf("error marshalling response: %v, value: %s", err, respBodyData)
	}

	data, err := json.Marshal(LogRequest{
		Method:   r.Method,
		Host:     r.Host,
		URL:      r.URL.String(),
		Headers:  r.Header,
		Body:     reqBody,
		Response: respBody,
	})
	if err != nil {
		return fmt.Errorf("error marshalling log request: %v", err)
	}

	p := filepath.Join(logDir, strconv.FormatInt(time.Now().UnixNano(), 10)+".json")

	return os.WriteFile(p, data, os.ModePerm)
}

func NewResponseLogger(w http.ResponseWriter) *ResponseLogger {
	return &ResponseLogger{
		ResponseWriter: w,
	}
}

// responseWriter wraps the http.ResponseWriter to capture the status code and size
type ResponseLogger struct {
	http.ResponseWriter
	statusCode int
	body       []byte
	size       int
}

// WriteHeader captures the status code
func (rw *ResponseLogger) WriteHeader(code int) {
	rw.statusCode = code
	// Not needed I guess
	//rw.ResponseWriter.WriteHeader(code)
}

// Write captures the size of the response body
func (rw *ResponseLogger) Write(b []byte) (int, error) {
	rw.body = append(rw.body, b...)

	n, err := rw.ResponseWriter.Write(b)
	rw.size += n
	return n, err
}

func (rw *ResponseLogger) Print() {
	fmt.Println("\n\nResponse:\n\n")
	fmt.Printf("Status Code: %d\n", rw.statusCode)
	fmt.Println("Body:\n")
	fmt.Println(string(rw.body))
	fmt.Println("\n\n")
}
