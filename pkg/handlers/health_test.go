package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/samber/lo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Create a new request to pass to our handler.
	body := map[string]any{
		"requestMetadata": map[string]any{
			"kmsRequestId":    "6fd27c22-d14a-4c83-8ad2-e4700eb78d87",
			"awsPrincipalArn": "AWS Internal",
			"KmsOperation":    "CreateCustomKeyStore",
		},
	}

	want := map[string]any{
		"ekmFleetDetails": []any{
			map[string]any{
				"healthStatus": "ACTIVE",
				"id":           "hsm-id-1",
				"model":        "Luna 5.0",
			},
		},
		"ekmVendor":         "Thales Group",
		"xksProxyFleetSize": float64(2),
		"xksProxyModel":     "Acme XKS Proxy 1.0",
		"xksProxyVendor":    "Acme Corp",
	}

	req := httptest.NewRequest("GET", "/yxks/kms/xks/v1/health", bytes.NewReader(lo.Must(json.Marshal(body))))

	// Create a ResponseRecorder to capture the response.
	rr := httptest.NewRecorder()

	// Call the handler, passing in the ResponseRecorder and the request.
	handler := http.HandlerFunc(HealthHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var resp map[string]any

	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("error unmarshalling response: %v", err)
	}

	if diff := cmp.Diff(resp, want); diff != "" {
		t.Errorf("differs: (-got +want)\n%s", diff)
	}
}

func TestHealth(t *testing.T) {
	type args struct {
		req GetHealthStatusRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *GetHealthStatusResponse
		wantErr bool
	}{
		{
			name: "Test Encrypt",
			args: args{
				req: GetHealthStatusRequest{
					RequestMetadata: HealthRequestMetadata{
						KmsOperation: "GetHealthStatus",
						KmsRequestId: "4112f4d6-db54-4af4-ae30-c55a22a8dfae",
					},
				},
			},
			want: &GetHealthStatusResponse{
				XksProxyFleetSize: 2,
				XksProxyVendor:    "Acme Corp",
				XksProxyModel:     "Acme XKS Proxy 1.0",
				EkmVendor:         "Thales Group",
				EkmFleetDetails: []EkmFleetDetail{
					{
						Id:           "hsm-id-1",
						Model:        "Luna 5.0",
						HealthStatus: "DEGRADED",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Health(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Encrypt() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
