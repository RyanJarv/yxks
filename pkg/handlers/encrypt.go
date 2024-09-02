package handlers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
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

// Encrypt function for the KMS XKS server
func Encrypt(id string, req EncryptRequest) (*EncryptResponse, error) {
	// Inline encryption key (this is just for demonstration, do not hardcode keys in production)
	encryptionKey := []byte("thisisa32byteencryptionkey!!!!!!") // 32 bytes for AES-256

	// Create AES cipher block
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	// Generate a random 12-byte initialization vector (IV) for AES-GCM
	iv := make([]byte, 12) // 12 bytes for AES-GCM nonce size
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Create AES-GCM cipher mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Encrypt the plaintext
	ciphertext := aesGCM.Seal(nil, iv, req.Plaintext, nil)

	// Compute the data integrity value using HMAC-SHA256
	h := hmac.New(sha256.New, encryptionKey)
	h.Write(ciphertext)
	dataIntegrityValue := h.Sum(nil)

	// Create the EncryptResponse
	response := &EncryptResponse{
		AuthenticationTag:            base64.StdEncoding.EncodeToString(ciphertext[len(ciphertext)-aesGCM.Overhead():]),
		Ciphertext:                   base64.StdEncoding.EncodeToString(ciphertext),
		CiphertextDataIntegrityValue: base64.StdEncoding.EncodeToString(dataIntegrityValue),
		CiphertextMetadata:           base64.StdEncoding.EncodeToString([]byte("key_version=1")), // Example metadata
		InitializationVector:         base64.StdEncoding.EncodeToString(iv),
	}

	return response, nil
}
