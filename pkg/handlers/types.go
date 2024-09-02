package handlers

// RequestMetadata represents the requestMetadata field in the EncryptRequest
// Example:
//
//	{
//	     "awsPrincipalArn": "arn:aws:iam::123456789012:user/Alice",
//	     "kmsKeyArn": "arn:aws:kms:us-east-2:123456789012:/key/1234abcd-12ab-34cd-56ef-1234567890ab",
//	     "kmsOperation": "Encrypt",
//	     "kmsRequestId": "4112f4d6-db54-4af4-ae30-c55a22a8dfae",
//	     "kmsViaService": "ebs"
//	 }
type RequestMetadata struct {
	AwsPrincipalArn string `json:"awsPrincipalArn"`
	KmsKeyArn       string `json:"kmsKeyArn"`
	KmsOperation    string `json:"kmsOperation"`
	KmsRequestId    string `json:"kmsRequestId"`
	KmsViaService   string `json:"kmsViaService"`
}

// EncryptRequest represents the request body for the Encrypt endpoint
// Example:
//
//	{
//	   "requestMetadata": ... ,
//	   "additionalAuthenticatedData": string (Base64 encoded), // optional
//	   "plaintext": string (Base64 encoded),
//	   "encryptionAlgorithm": string,
//	   "ciphertextDataIntegrityValueAlgorithm": string // optional
//	 }
type EncryptRequest struct {
	RequestMetadata                       RequestMetadata `json:"requestMetadata"`
	AdditionalAuthenticatedData           string          `json:"additionalAuthenticatedData"`
	Plaintext                             []byte          `json:"plaintext"`
	EncryptionAlgorithm                   string          `json:"encryptionAlgorithm"`
	CiphertextDataIntegrityValueAlgorithm string          `json:"ciphertextDataIntegrityValueAlgorithm"`
}

// EncryptResponse represents the response body for the Encrypt endpoint
// Example:
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
