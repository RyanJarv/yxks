package handlers

import (
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		id  string
		req EncryptRequest
	}
	tests := []struct {
		name    string
		args    args
		want    EncryptResponse
		wantErr bool
	}{
		{
			name: "Test Encrypt",
			args: args{
				id: "1234abcd-12ab-34cd-56ef-1234567890ab",
				req: EncryptRequest{
					AuthenticationTag:            "vBxN2ncH1oEkR8WVXpmyYQ==",
					Ciphertext:                   "ghxkK1txeDNn3q8Y",
					CiphertextDataIntegrityValue: "qHA/ImC9h5HsLRXqCyPmWgYx7tzyoTplzILbP0fPXsc=",
					CiphertextMetadata:           "a2V5X3ZlcnNpb249MQ==",
				},
			},
			want: EncryptResponse{
				AuthenticationTag:            "vBxN2ncH1oEkR8WVXpmyYQ==",
				Ciphertext:                   "ghxkK1txeDNn3q8Y",
				CiphertextDataIntegrityValue: "qHA/ImC9h5HsLRXqCyPmWgYx7tzyoTplzILbP0fPXsc=",
				CiphertextMetadata:           "a2V5X3ZlcnNpb249MQ==",
				InitializationVector:         "HMrlRw85cAJUd5Ax",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
