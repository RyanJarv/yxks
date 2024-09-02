# YXKS


## Generate shared secret

openssl rand 32 | base64


## Paths

URI: https://<server>[/<path-prefix>]/kms/xks/v1/keys/<externalKeyId>/metadata
URI: https://<server>[/<path-prefix>]/kms/xks/v1/keys/<externalKeyId>/encrypt
URI: https://<server>[/<path-prefix>]/kms/xks/v1/keys/<externalKeyId>/decrypt
URI: https://<server>[/<path-prefix>]/kms/xks/v1/health



## GetKeyMetadata -- xks_proxy_api_spec.md

HTTP Method: POST
URI: https://<server>[/<path-prefix>]/kms/xks/v1/keys/<externalKeyId>/metadata

### Request

```
{
    "requestMetadata": {
        "awsPrincipalArn": "arn:aws:iam::123456789012:user/Alice",
        "kmsOperation": "CreateKey",
        "kmsRequestId": "4112f4d6-db54-4af4-ae30-c55a22a8dfae"
    }
}
```


### Response

```
{
    "requestMetadata": {
        "awsPrincipalArn": "arn:aws:iam::123456789012:user/Alice",
        "kmsKeyArn": "arn:aws:kms:us-east-2:123456789012:/key/1234abcd-12ab-34cd-56ef-1234567890ab",
        "kmsOperation": "Encrypt",
        "kmsRequestId": "4112f4d6-db54-4af4-ae30-c55a22a8dfae",
        "kmsViaService": "ebs"
    }, 
    "additionalAuthenticatedData": "cHJvamVjdD1uaWxlLGRlcGFydG1lbnQ9bWFya2V0aW5n",
    "plaintext": "SGVsbG8gV29ybGQh",
    "encryptionAlgorithm": "AES_GCM",
    "ciphertextDataIntegrityValueAlgorithm": "SHA_256"
}
```


## Encrypt -- xks_proxy_api_spec.md

HTTP Method: POST
URI: https://<server>[/<path-prefix>]/kms/xks/v1/keys/<externalKeyId>/encrypt

### Request

```
{
    "requestMetadata": {
        "awsPrincipalArn": "arn:aws:iam::123456789012:user/Alice",
        "kmsKeyArn": "arn:aws:kms:us-east-2:123456789012:/key/1234abcd-12ab-34cd-56ef-1234567890ab",
        "kmsOperation": "Encrypt",
        "kmsRequestId": "4112f4d6-db54-4af4-ae30-c55a22a8dfae",
        "kmsViaService": "ebs"
    }, 
    "additionalAuthenticatedData": "cHJvamVjdD1uaWxlLGRlcGFydG1lbnQ9bWFya2V0aW5n",
    "plaintext": "SGVsbG8gV29ybGQh",
    "encryptionAlgorithm": "AES_GCM",
    "ciphertextDataIntegrityValueAlgorithm": "SHA_256"
}
```


### Response

```
{
    "authenticationTag": "vBxN2ncH1oEkR8WVXpmyYQ==",
    "ciphertext": "ghxkK1txeDNn3q8Y",
    "ciphertextDataIntegrityValue": "qHA/ImC9h5HsLRXqCyPmWgYx7tzyoTplzILbP0fPXsc=",
    "ciphertextMetadata": "a2V5X3ZlcnNpb249MQ==",
    "initializationVector": "HMrlRw85cAJUd5Ax"
}
```


## Decrypt -- xks_proxy_api_spec.md

HTTP Method: POST
URI: https://<server>[/<path-prefix>]/kms/xks/v1/keys/<externalKeyId>/decrypt

### Request

```
{
    "requestMetadata": { 
        "awsPrincipalArn": "arn:aws:iam::123456789012:user/Alice",
        "kmsKeyArn": "arn:aws:kms:us-east-2:123456789012:/key/1234abcd-12ab-34cd-56ef-1234567890ab",
        "kmsOperation": "Decrypt",
        "kmsRequestId": "5112f4d6-db54-4af4-ae30-c55a22a8dfae",
        "kmsViaService": "ebs"
    },
    "additionalAuthenticatedData": "cHJvamVjdD1uaWxlLGRlcGFydG1lbnQ9bWFya2V0aW5n",
    "encryptionAlgorithm": "AES_GCM",
    "ciphertext": "ghxkK1txeDNn3q8Y",
    "ciphertextMetadata": "a2V5X3ZlcnNpb249MQ==",
    "initializationVector": "HMrlRw85cAJUd5Ax",
    "authenticationTag": "vBxN2ncH1oEkR8WVXpmyYQ=="
}
```

### Response

```
{
    "plaintext": "SGVsbG8gV29ybGQh"
}
```



## GetHealthStatus -- xks_proxy_api_spec.md


HTTP Method: POST
URI: https://<server>[/<path-prefix>]/kms/xks/v1/health


### Request

```
{
    "requestMetadata": {
        "kmsRequestId": "4112f4d6-db54-4af4-ae30-c55a22a8dfae",
        "kmsOperation": "CreateCustomKeyStore"
    }
}
```

### Response

```
{
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
}
```
