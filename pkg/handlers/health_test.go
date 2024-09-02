package handlers

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

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
