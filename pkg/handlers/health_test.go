package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
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
			name: "Test Health",
			args: args{
				w: httptest.NewRecorder(),
				req: httptest.NewRequest(http.MethodPost, "/kms/xks/v1/health", bytes.NewBuffer([]byte(`{
    "requestMetadata": {
        "kmsRequestId": "4112f4d6-db54-4af4-ae30-c55a22a8dfae",
        "kmsOperation": "CreateCustomKeyStore"
    }
}`))),
			},
			wantStatus: http.StatusOK,
			wantResp: `{
    "xksProxyFleetSize": 2,
    "xksProxyVendor": "Acme Corp",
    "xksProxyModel": "Acme XKS Proxy 1.0",
    "ekmVendor": "Thales Group",
    "ekmFleetDetails": [
        {
            "id": "hsm-id-1",
            "model": "Luna 5.0",
            "healthStatus": "DEGRADED"
        },
        {
            "id": "hsm-id-2",
            "model": "Luna 5.1",
            "healthStatus": "ACTIVE"
        }
    ]
}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Health(tt.args.w, tt.args.req)
			if tt.args.w.(*httptest.ResponseRecorder).Code != tt.wantStatus {
				t.Errorf("Health Status Code: got %v, want %v", tt.args.w.(*httptest.ResponseRecorder).Code, tt.wantStatus)
			}

			if tt.args.w.(*httptest.ResponseRecorder).Body.String() != tt.wantResp {
				t.Errorf("Health Response: got %v, want %v", tt.args.w.(*httptest.ResponseRecorder).Body.String(), tt.wantResp)
			}
		})
	}
}
