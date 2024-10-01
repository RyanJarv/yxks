package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/samber/lo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDecryptHandler(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantResp   string
	}{
		{
			name: "Test HealthHandler",
			args: args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/kms/xks/v1/health", bytes.NewBuffer([]byte(`{"requestMetadata":{"kmsRequestId":"81cd476a-cbbf-45f2-a535-8fe167d13336","kmsKeyArn":"arn:aws:kms:us-east-2:123456789012:key/7d2bfabe-63ea-4c05-80cb-89300aa53151","kmsViaService":"s3.us-east-2.amazonaws.com","awsPrincipalArn":"ecr.amazonaws.com","kmsOperation":"GenerateDataKey"},"initializationVector":"rZmHwyP68KYI6gZj","ciphertext":"3w7MZCMJQCdUY0rHPzD9ryn0gjc6CoHdggApeE8Jh0BE8KgsY/C0vf1u3zatacIKL4T3CD6pmbsDOMIwXEebrThNJkv/P3rXSt2Oa9kKJFoyEwejFncDz+4QfmRKgpUVnZ4PuY/wvOjR2TE9+qDLnanrS7iz3sL/c04S4+90AanfSlRqmSXGMJx51KtUmwE18mwTXRS75vjamsMxGzQ1NjAU","additionalAuthenticatedData":"XQZhAAACAAthd3M6ZWNyOmFybgA2YXJuOmF3czplY3I6dXMtZWFzdC0yOjQwMzU4MzMxMjI4MjpyZXBvc2l0b3J5L2ttcy10ZXN0AAphd3M6czM6YXJuAI9hcm46YXdzOnMzOjo6cHJvZC11cy1lYXN0LTItc3RhcnBvcnQtbGF5ZXItYnVja2V0LzUyMTk1MC00MDM1ODMzMTIyODItYWVjOTIyZWQtN2JiMC1hMmEwLTM5MjYtOTU3MzRjMGM4MWQ5LzNlNTg0YzM2LWYzYjEtNDE0My1iYmVkLTAzNDc1ZTcwYTU0OYELLEFm+UImr95VSwiqH7KM+SaSek2BU4tnLR2+49PC","encryptionAlgorithm":"AES_GCM","authenticationTag":"PF46NkpyXsAbkuky1HQJ3Q=="}`))),
			},
			wantStatus: http.StatusOK,
			wantResp:   `{"plaintext": "MIGfBgkqhkiG9w0BBwaggZEwgY4CAQAwgYgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMV5A2SI8UznSVQSyeAgEQgFtc8EFiBS88UsEk3yuHkbRDfpBwXCUjLdymq25i1pexpnvgUuUcVuu37lpavCqcjs7E0y1lJJhucTnL6M8T2zXSHk6PvhtpJz6nf69EByyay+iHR/1Dg5EkeWCX"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DecryptHandler(tt.args.w, tt.args.req)
			if tt.args.w.(*httptest.ResponseRecorder).Code != tt.wantStatus {
				t.Errorf("HealthHandler Status Code: got %v, want %v", tt.args.w.(*httptest.ResponseRecorder).Code, tt.wantStatus)
			}

			var want = map[string]string{}
			lo.Must0(json.Unmarshal([]byte(tt.wantResp), &want))

			body := tt.args.w.(*httptest.ResponseRecorder).Body.Bytes()

			var got = map[string]string{}
			lo.Must0(json.Unmarshal(body, &got))
			fmt.Println(got)

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("HealthHandler Response: -want +got %s", diff)
			}
		})
	}
}
