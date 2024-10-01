package handlers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ryanjarv/yxks/pkg/utils"
	"github.com/samber/lo"
	"log"
	"net/http"
)

type DecryptRequestMetadata struct {
	AwsPrincipalArn string `json:"awsPrincipalArn"`
	KmsKeyArn       string `json:"kmsKeyArn"`
	KmsOperation    string `json:"kmsOperation"`
	KmsRequestId    string `json:"kmsRequestId"`
	KmsViaService   string `json:"kmsViaService"`
}

type DecryptRequest struct {
	DecryptRequestMetadata      DecryptRequestMetadata `json:"requestMetadata"`
	AdditionalAuthenticatedData string                 `json:"additionalAuthenticatedData"`
	EncryptionAlgorithm         string                 `json:"encryptionAlgorithm"`
	Ciphertext                  string                 `json:"ciphertext"`
	CiphertextMetadata          string                 `json:"ciphertextMetadata"`
	InitializationVector        string                 `json:"initializationVector"`
	AuthenticationTag           string                 `json:"authenticationTag"`
}

type DecryptResponse struct {
	Plaintext string `json:"plaintext"`
}

// DecryptHandler endpoint
// URI: /kms/xks/v1/keys/{externalKeyId}/decrypt
func DecryptHandler(w http.ResponseWriter, req *http.Request) {
	extKeyId := utils.GetExternalKeyId(req)

	utils.H[DecryptRequest](w, req, func(t DecryptRequest) (any, error) {
		log.Printf("[DEBUG] Decrypting with key: %s, input: %v\n", extKeyId, utils.TryJson(t))

		response, err := Decrypt(extKeyId, t)
		if err != nil {
			log.Printf("[ERROR] encrypting: %v", err)
			return nil, err
		}

		return response, nil
	})
}

// Decrypt function for the KMS XKS server
func Decrypt(id string, req DecryptRequest) (*DecryptResponse, error) {
	if req.AdditionalAuthenticatedData != "" {
		decodeString, err := base64.StdEncoding.DecodeString(req.AdditionalAuthenticatedData)
		if err != nil {
			return nil, fmt.Errorf("error decoding additional authenticated data: %v", err)
		}

		aad, leftover, err := ParseAAD(decodeString)
		if err != nil {
			return nil, fmt.Errorf("error parsing AAD: %v", err)
		}
		fmt.Println("Parsed additional authenticated data:", string(lo.Must(json.Marshal(aad))))
		if len(leftover) > 0 {
			fmt.Printf("Leftover additional authenticated data (%d bytes): %x\n", len(leftover), leftover)
		}

	}

	// Create AES cipher block
	block, err := aes.NewCipher(EncryptionKey)
	if err != nil {
		return nil, err
	}

	// Decode Base64-encoded inputs
	ciphertext, err := base64.StdEncoding.DecodeString(req.Ciphertext)
	if err != nil {
		return nil, err
	}

	iv, err := base64.StdEncoding.DecodeString(req.InitializationVector)
	if err != nil {
		return nil, err
	}

	authTag, err := base64.StdEncoding.DecodeString(req.AuthenticationTag)
	if err != nil {
		return nil, err
	}

	// NOTE: We're skipping AAD for now
	//var additionalAuthenticatedData []byte
	//if req.AdditionalAuthenticatedData != "" {
	//	additionalAuthenticatedData, err = base64.StdEncoding.DecodeString(req.AdditionalAuthenticatedData)
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	// Create AES-GCM cipher mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Decrypt the ciphertext
	plaintext, err := aesGCM.Open(nil, iv, append(ciphertext, authTag...), nil)
	if err != nil {
		return nil, err
	}

	ParsePlaintextData(plaintext)

	response := &DecryptResponse{
		Plaintext: base64.StdEncoding.EncodeToString(plaintext),
	}

	return response, nil
}
