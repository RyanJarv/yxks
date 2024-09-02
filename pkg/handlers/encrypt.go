package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ryanjarv/yxks/pkg/utils"
	"io"
	"net/http"
)

// Encrypt endpoint
// URI: /kms/xks/v1/keys/{externalKeyId}/Encrypt
func EncryptHandler(w http.ResponseWriter, req *http.Request) {
	err := encryptHandlerErr(w, req)
	if err != nil {
		panic(err)
	}
}

func encryptHandlerErr(w http.ResponseWriter, req *http.Request) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("error reading request body: %v", err)
	}

	var encryptReq EncryptRequest
	if err := json.Unmarshal(body, &encryptReq); err != nil {
		return fmt.Errorf("error unmarshalling request body: %v", err)
	}

	extKeyId := utils.GetExternalKeyId(req)

	response, err := Encrypt(extKeyId, encryptReq)
	if err != nil {
		return fmt.Errorf("error encrypting: %v", err)
	}

	marshal, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("error marshalling response: %v", err)
	}

	if _, err := w.Write(marshal); err != nil {
		return fmt.Errorf("error writing response: %v", err)
	}

	return nil
}

// EncryptRequest represents the request body for the Encrypt endpoint
// Exmaple JSON request body:
//
//	 {
//	    "authenticationTag": "vBxN2ncH1oEkR8WVXpmyYQ==",
//	    "ciphertext": "ghxkK1txeDNn3q8Y",
//	    "ciphertextDataIntegrityValue": "qHA/ImC9h5HsLRXqCyPmWgYx7tzyoTplzILbP0fPXsc=",
//	    "ciphertextMetadata": "a2V5X3ZlcnNpb249MQ==",
//	    "initializationVector": "HMrlRw85cAJUd5Ax"
//	}
type EncryptRequest struct {
	AuthenticationTag            string `json:"authenticationTag"`
	Ciphertext                   string `json:"ciphertext"`
	CiphertextDataIntegrityValue string `json:"ciphertextDataIntegrityValue"`
	CiphertextMetadata           string `json:"ciphertextMetadata"`
}

// EncryptResponse represents the response body for the Encrypt endpoint
// Example JSON response body:
//
//	{
//	  "authenticationTag": "vBxN2ncH1oEkR8WVXpmyYQ==",
//	  "ciphertext": "ghxkK1txeDNn3q8Y",
//	  "ciphertextDataIntegrityValue": "qHA/ImC9h5HsLRXqCyPmWgYx7tzyoTplzILbP0fPXsc=",
//	  "ciphertextMetadata": "a2V5X3ZlcnNpb249MQ==",
//	  "initializationVector": "HMrlRw85cAJUd5Ax"
//	}
type EncryptResponse struct {
	AuthenticationTag            string `json:"authenticationTag"`
	Ciphertext                   string `json:"ciphertext"`
	CiphertextDataIntegrityValue string `json:"ciphertextDataIntegrityValue"`
	CiphertextMetadata           string `json:"ciphertextMetadata"`
	InitializationVector         string `json:"initializationVector"`
}

func Encrypt(id string, req EncryptRequest) (EncryptResponse, error) {
	return EncryptResponse{
		AuthenticationTag:            "vBxN2ncH1oEkR8WVXpmyYQ==",
		Ciphertext:                   "ghxkK1txeDNn3q8Y",
		CiphertextDataIntegrityValue: "qHA/ImC9h5HsLRXqCyPmWgYx7tzyoTplzILbP0fPXsc=",
		CiphertextMetadata:           "a2V5X3ZlcnNpb249MQ==",
		InitializationVector:         "HMrlRw85cAJUd5Ax",
	}, nil
}
