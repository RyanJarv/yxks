package handlers

// EncryptRequestMetadata represents the requestMetadata field in the EncryptRequest
// Example:
//
//	{
//	     "awsPrincipalArn": "arn:aws:iam::123456789012:user/Alice",
//	     "kmsKeyArn": "arn:aws:kms:us-east-2:123456789012:/key/1234abcd-12ab-34cd-56ef-1234567890ab",
//	     "kmsOperation": "Encrypt",
//	     "kmsRequestId": "4112f4d6-db54-4af4-ae30-c55a22a8dfae",
//	     "kmsViaService": "ebs"
//	 }
type EncryptRequestMetadata struct {
	AwsPrincipalArn string `json:"awsPrincipalArn"`
	KmsKeyArn       string `json:"kmsKeyArn"`
	KmsOperation    string `json:"kmsOperation"`
	KmsRequestId    string `json:"kmsRequestId"`
	KmsViaService   string `json:"kmsViaService"`
}

// EncryptRequest represents the request body for the Encrypt endpoint
// Example:
//
//		{
//	   "requestMetadata": ...,
//	   "additionalAuthenticatedData": "cHJvamVjdD1uaWxlLGRlcGFydG1lbnQ9bWFya2V0aW5n",
//	   "plaintext": "SGVsbG8gV29ybGQh",
//	   "encryptionAlgorithm": "AES_GCM",
//	   "ciphertextDataIntegrityValueAlgorithm": "SHA_256"
//	 }
type EncryptRequest struct {
	RequestMetadata                       EncryptRequestMetadata `json:"requestMetadata"`
	AdditionalAuthenticatedData           string                 `json:"additionalAuthenticatedData"`
	Plaintext                             []byte                 `json:"plaintext"`
	EncryptionAlgorithm                   string                 `json:"encryptionAlgorithm"`
	CiphertextDataIntegrityValueAlgorithm string                 `json:"ciphertextDataIntegrityValueAlgorithm"`
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

// HealthRequestMetadata represents the requestMetadata field in the GetHealthStatusRequest
// Example:
//
//	{
//	  "kmsRequestId": "4112f4d6-db54-4af4-ae30-c55a22a8dfae",
//	  "kmsOperation": "CreateCustomKeyStore"
//	}
type HealthRequestMetadata struct {
	KmsRequestId string `json:"kmsRequestId"`
	KmsOperation string `json:"kmsOperation"`
}

// GetHealthStatusRequest represents the request body for the GetHealthStatus endpoint
// Example:
//
//	{
//	   "requestMetadata": {
//	       "kmsRequestId": "4112f4d6-db54-4af4-ae30-c55a22a8dfae",
//	       "kmsOperation": "CreateCustomKeyStore"
//	   }
//	}
type GetHealthStatusRequest struct {
	RequestMetadata HealthRequestMetadata `json:"requestMetadata"`
}

// EkmFleetDetail represents the EkmFleetDetail field in the GetHealthStatusResponse
// Example:
//
//	{
//	    "id": "hsm-id-1",
//	    "model": "Luna 5.0",
//	    "healthStatus": "DEGRADED"
//	}
type EkmFleetDetail struct {
	Id           string `json:"id"`
	Model        string `json:"model"`
	HealthStatus string `json:"healthStatus"`
}

// GetHealthStatusResponse represents the response body for the GetHealthStatus endpoint
// Example:
//
//	{
//	   "xksProxyFleetSize": 2,
//	   "xksProxyVendor": "Acme Corp",
//	   "xksProxyModel": "Acme XKS Proxy 1.0",
//	   "ekmVendor": "Thales Group",
//	   "ekmFleetDetails": [
//	       {
//	           "id": "hsm-id-1",
//	           "model": "Luna 5.0",
//	           "healthStatus": "DEGRADED"
//	       },
//	       {
//	           "id": "hsm-id-2",
//	           "model": "Luna 5.1",
//	           "healthStatus": "ACTIVE"
//	       }
//	   ]
//	 }
type GetHealthStatusResponse struct {
	XksProxyFleetSize int              `json:"xksProxyFleetSize"`
	XksProxyVendor    string           `json:"xksProxyVendor"`
	XksProxyModel     string           `json:"xksProxyModel"`
	EkmVendor         string           `json:"ekmVendor"`
	EkmFleetDetails   []EkmFleetDetail `json:"ekmFleetDetails"`
}
