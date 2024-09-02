package handlers

import (
	"bytes"
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
				w: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/kms/xks/v1/health", bytes.NewBuffer([]byte(`{
    "requestMetadata": { 
        "awsPrincipalArn": "arn:aws:iam::123456789012:user/Alice",
        "kmsKeyArn": "arn:aws:kms:us-east-2:123456789012:/key/1234abcd-12ab-34cd-56ef-1234567890ab",
        "kmsOperation": "DecryptHandler",
        "kmsRequestId": "5112f4d6-db54-4af4-ae30-c55a22a8dfae",
        "kmsViaService": "ebs"
    },
    "additionalAuthenticatedData": "cHJvamVjdD1uaWxlLGRlcGFydG1lbnQ9bWFya2V0aW5n",
    "encryptionAlgorithm": "AES_GCM",
    "ciphertext": "ghxkK1txeDNn3q8Y",
    "ciphertextMetadata": "a2V5X3ZlcnNpb249MQ==",
    "initializationVector": "HMrlRw85cAJUd5Ax",
    "authenticationTag": "vBxN2ncH1oEkR8WVXpmyYQ=="
}`))),
			},
			wantStatus: http.StatusOK,
			wantResp: `{
    "plaintext": "SGVsbG8gV29ybGQh"
}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DecryptHandler(tt.args.w, tt.args.req)
			if tt.args.w.(*httptest.ResponseRecorder).Code != tt.wantStatus {
				t.Errorf("HealthHandler Status Code: got %v, want %v", tt.args.w.(*httptest.ResponseRecorder).Code, tt.wantStatus)
			}

			if tt.args.w.(*httptest.ResponseRecorder).Body.String() != tt.wantResp {
				t.Errorf("HealthHandler Response: got %v, want %v", tt.args.w.(*httptest.ResponseRecorder).Body.String(), tt.wantResp)
			}
		})
	}
}
