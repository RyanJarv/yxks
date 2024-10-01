package handlers

import (
	"github.com/ryanjarv/yxks/pkg/utils"
	"log"
	"net/http"
)

// GetKeyMetadataHandler endpoint
// URI: /kms/xks/v1/keys/{externalKeyId}/metadata
func GetKeyMetadataHandler(w http.ResponseWriter, req *http.Request) {
	extKeyId := utils.GetExternalKeyId(req)

	utils.H[KeyMetadataRequest](w, req, func(t KeyMetadataRequest) (any, error) {
		log.Printf("[DEBUG] metdata with key: %s, input: %v\n", extKeyId, utils.TryJson(t))

		return KeyMetadataResponse{
			KeySpec: "AES_256",
			KeyUsage: []string{
				"ENCRYPT",
				"DECRYPT",
				"DERIVE",
				"SIGN",
				"VERIFY",
				"WRAP",
				"UNWRAP",
			},
			KeyStatus: "ENABLED",
		}, nil
	})
}

type RequestMetadata struct {
	AwsPrincipalArn string `json:"awsPrincipalArn"`
	AwsSourceVpc    string `json:"awsSourceVpc"`
	AwsSourceVpce   string `json:"awsSourceVpce"`
	KmsOperation    string `json:"KmsOperation"`
	KmsRequestId    string `json:"kmsRequestId"`
}

type KeyMetadataRequest struct {
	RequestMetadata RequestMetadata `json:"requestMetadata"`
}

type KeyMetadataResponse struct {
	KeySpec   string   `json:"keySpec"`
	KeyUsage  []string `json:"keyUsage"`
	KeyStatus string   `json:"keyStatus"`
}
