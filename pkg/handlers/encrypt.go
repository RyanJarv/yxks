package handlers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/asn1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/ryanjarv/yxks/pkg/utils"
	"github.com/samber/lo"
	"io"
	"log"
	"net/http"
)

// EncryptionKey Inline encryption key (this is just for demonstration; do not hardcode keys in production)
var EncryptionKey = []byte("thisisa32byteencryptionkey!!!!!!") // 32 bytes for AES-256

// EncryptHandler endpoint
// URI: /kms/xks/v1/keys/{externalKeyId}/Encrypt
func EncryptHandler(w http.ResponseWriter, req *http.Request) {
	extKeyId := utils.GetExternalKeyId(req)

	utils.H[EncryptRequest](w, req, func(t EncryptRequest) (any, error) {
		log.Printf("[DEBUG] Encrypting plaintext for key: %s, input: %s\n", extKeyId, utils.TryJson(t))

		response, err := Encrypt(extKeyId, t)
		if err != nil {
			log.Printf("[ERROR] error encrypting: %v", err)
			return nil, err
		}

		return response, nil
	})

}

func Encrypt(id string, req EncryptRequest) (*EncryptResponse, error) {
	if req.AdditionalAuthenticatedData != "" {
		decodeString, err := base64.StdEncoding.DecodeString(req.AdditionalAuthenticatedData)
		if err != nil {
			return nil, fmt.Errorf("error decoding additional authenticated data: %v", err)
		}

		aad, _, err := ParseAAD(decodeString)
		if err != nil {
			return nil, fmt.Errorf("error parsing AAD: %v", err)
		}
		fmt.Println("additional authenticated data:", string(lo.Must(json.Marshal(aad))))

	}

	// Create AES cipher block
	block, err := aes.NewCipher(EncryptionKey)
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

	// Decode the Base64-encoded plaintext
	decodedPlaintext, err := base64.StdEncoding.DecodeString(req.Plaintext) // Convert []byte to string
	if err != nil {
		return nil, err
	}

	ParsePlaintextData(decodedPlaintext)

	// Encrypt the decoded plaintext
	ciphertext := aesGCM.Seal(nil, iv, decodedPlaintext, nil)

	// Extract the authentication tag (last part of the ciphertext)
	authTag := ciphertext[len(ciphertext)-aesGCM.Overhead():]
	ciphertextWithoutAuthTag := ciphertext[:len(ciphertext)-aesGCM.Overhead()]

	// Create the EncryptResponse
	resp := &EncryptResponse{
		AuthenticationTag:    base64.StdEncoding.EncodeToString(authTag),
		Ciphertext:           base64.StdEncoding.EncodeToString(ciphertextWithoutAuthTag),
		InitializationVector: base64.StdEncoding.EncodeToString(iv),
	}

	return resp, nil
}

func ParsePlaintextData(derData []byte) {
	// Unmarshal the ASN.1 structure based on the PKCS#7 EncryptedData definition
	var pkcs7Data PKCS7EncryptedData
	if _, err := asn1.Unmarshal(derData, &pkcs7Data); err != nil {
		log.Fatalf("Failed to unmarshal ASN.1 data: %v", err)
	}

	// Parse the AES-GCM Parameters
	var aesGCMParams AESGCMParameters
	if _, err := asn1.Unmarshal(pkcs7Data.Content.EncryptedContentInfo.ContentEncryptionAlgorithm.Parameters.FullBytes, &aesGCMParams); err != nil {
		log.Fatalf("Failed to unmarshal AES-GCM parameters: %v", err)
	}

	// Output the parsed data
	fmt.Printf("KMS pkcs7 plaintext:\n")
	fmt.Printf("  Content Type: %s\n", pkcs7Data.ContentType)
	fmt.Printf("  Version: %d\n", pkcs7Data.Content.Version)
	fmt.Printf("  Encrypted Content Info:\n")
	fmt.Printf("  	Content Type: %s\n", pkcs7Data.Content.EncryptedContentInfo.ContentType)
	fmt.Printf("  	Encryption Algorithm: %s\n", pkcs7Data.Content.EncryptedContentInfo.ContentEncryptionAlgorithm.Algorithm)
	fmt.Printf("  	Nonce (IV): %x\n", aesGCMParams.Nonce)
	fmt.Printf("  	ICV Length: %d\n", aesGCMParams.ICVLen)
	fmt.Printf("  	Encrypted Content: %x\n", string(pkcs7Data.Content.EncryptedContentInfo.EncryptedContent))
}

type PKCS7EncryptedData struct {
	ContentType asn1.ObjectIdentifier
	Content     EncryptedDataContent `asn1:"explicit,tag:0"`
}

type EncryptedDataContent struct {
	Version              int
	EncryptedContentInfo EncryptedContentInfo
}

type EncryptedContentInfo struct {
	ContentType                asn1.ObjectIdentifier
	ContentEncryptionAlgorithm AlgorithmIdentifier
	EncryptedContent           []byte `asn1:"optional,tag:0,implicit"`
}

type AlgorithmIdentifier struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.RawValue `asn1:"optional"`
}

type AESGCMParameters struct {
	Nonce  []byte
	ICVLen int `asn1:"optional,default:12"`
}

func ParseAAD(data []byte) (map[string]string, []byte, error) {
	reader := bytes.NewReader(data)
	var err error

	// Read the magic number (4 bytes)
	magicNumber := make([]byte, 4)
	if _, err = reader.Read(magicNumber); err != nil {
		return nil, nil, err
	}
	// Optionally verify the magic number
	if !bytes.Equal(magicNumber, []byte{0x5d, 0x06, 0x61, 0x00}) {
		return nil, nil, fmt.Errorf("invalid magic number: %x", magicNumber)
	}

	// Read the number of key-value pairs (2 bytes, big-endian)
	var numPairs uint16
	if err = binary.Read(reader, binary.BigEndian, &numPairs); err != nil {
		return nil, nil, err
	}

	pairs := make(map[string]string)
	for i := 0; i < int(numPairs); i++ {
		// Read key length (2 bytes)
		var keyLength uint16
		if err = binary.Read(reader, binary.BigEndian, &keyLength); err != nil {
			return nil, nil, err
		}

		// Read key
		keyBytes := make([]byte, keyLength)
		if _, err = reader.Read(keyBytes); err != nil {
			return nil, nil, err
		}
		key := string(keyBytes)

		// Read value length (2 bytes)
		var valueLength uint16
		if err = binary.Read(reader, binary.BigEndian, &valueLength); err != nil {
			return nil, nil, err
		}

		// Read value
		valueBytes := make([]byte, valueLength)
		if _, err = reader.Read(valueBytes); err != nil {
			return nil, nil, err
		}
		value := string(valueBytes)

		// Store the key-value pair
		pairs[key] = value
	}

	// ChatGPT's take on this leftover data:
	//
	//  Possible Interpretation:
	//
	//  The fact that each of these data samples is exactly 32 bytes (256 bits) suggests that they may represent
	//  cryptographic material such as:
	//
	//  SHA-256 Hashes: The output size of the SHA-256 hashing algorithm is 32 bytes.
	//  HMAC-SHA256 Authentication Tags: An HMAC using SHA-256 produces a 32-byte output.
	//  AES-GCM Authentication Tags: Typically, AES-GCM uses a 16-byte (128-bit) authentication tag by default, but it
	//  								can be configured to use different tag lengths, although 32 bytes is uncommon.
	//
	//  Given that this data follows the key-value pairs (the additional authenticated data, or AAD) and is associated
	//  with AWS KMS operations involving an XKS proxy, it's likely that this leftover data is an authentication tag
	//  or message authentication code (MAC) used to verify the integrity and authenticity of the data.
	//
	leftover, err := io.ReadAll(reader)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading remaining bytes: %v", err)
	}

	return pairs, leftover, nil
}
