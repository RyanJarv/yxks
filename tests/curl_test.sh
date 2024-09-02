#!/bin/bash

# XKS proxy parameters
XKS_PROXY_URI_ENDPOINT=http://localhost:8080
XKS_PROXY_URI_PATH=/kms/xks/v1
ACCESS_KEY_ID=HGIE2BUAFNDMBCMSQVWLGCEXAMPLE
SECRET_ACCESS_KEY=VU6W2byoiBrhDqNGbx703oL0syCAmhlQgwNUT123456
XKS_KEY_ID=b42c7fbf-de61-441c-ae24-123456789012



# See the -w option in curl
cat > curl_metrics.txt << EOF
             http_version:  %{http_version}\n
                      url:  %{url}\n
        request_body_size:  %{size_upload}\n
     response_header_size:  %{size_header}\n
       response_body_size:  %{size_download}\n
            response_code:  %{response_code}\n
 tls_verification(0=true):  %{ssl_verify_result}\n
      start_to_dns_lookup:  %{time_namelookup}s\n
     start_to_tcp_connect:  %{time_connect}s\n
   start_to_tls_handshake:  %{time_appconnect}s\n
             start_to_end:  %{time_total}s\n
EOF


# Sending an Encrypt XKS API request using curl
XKS_ENCRYPT_URI=${XKS_PROXY_URI_ENDPOINT}${XKS_PROXY_URI_PATH}/keys/${XKS_KEY_ID}/encrypt

cat > xks_encrypt_request.json << EOF
{
    "requestMetadata": {
        "awsPrincipalArn": "arn:aws:iam::123456789012:user/Alice",
        "kmsKeyArn": "arn:aws:kms:us-east-2:123456789012:/key/1234abcd-12ab-34cd-56ef-1234567890ab",
        "kmsOperation": "GenerateDataKey",
        "kmsRequestId": "4112f4d6-db54-4af4-ae30-c55a22a8dfae"
    },
    "plaintext": "SGVsbG8gV29ybGQh",
    "encryptionAlgorithm": "AES_GCM"
}
EOF

# Send Encrypt API request, view response and timing data
echo "Making Encrypt XKS API call ..."
curl --aws-sigv4  "aws:amz:us-east-1:kms-xks-proxy" \
    --user ${ACCESS_KEY_ID}:${SECRET_ACCESS_KEY} \
    -X POST -H "Content-Type:application/json" \
    --data @xks_encrypt_request.json ${XKS_ENCRYPT_URI}

echo
echo
echo "Collecting Encrypt metrics ..."
curl --aws-sigv4  "aws:amz:us-east-1:kms-xks-proxy" \
    --user ${ACCESS_KEY_ID}:${SECRET_ACCESS_KEY} \
    -X POST -H "Content-Type:application/json" \
    --data @xks_encrypt_request.json \
    -w "@curl_metrics.txt" -o encrypt.out -s \
    ${XKS_ENCRYPT_URI}



# Sending a GetHealthStatus XKS API request using curl
XKS_GET_HEALTH_STATUS_URI=${XKS_PROXY_URI_ENDPOINT}${XKS_PROXY_URI_PATH}/health

cat > xks_get_health_status_request.json << EOF
{
    "requestMetadata": {
        "kmsOperation": "KmsHealthCheck",
        "kmsRequestId": "1124f4d6-db54-4af4-ae30-c55a22a8abcd"
    }
}
EOF

# Send GetHealthStatus API request, view response and timing data
echo "Making GetHealthStatus XKS API call ..."
curl --aws-sigv4  "aws:amz:us-east-1:kms-xks-proxy" \
  --user ${ACCESS_KEY_ID}:${SECRET_ACCESS_KEY} \
    -X POST -H "Content-Type:application/json" \
    --data @xks_get_health_status_request.json  ${XKS_GET_HEALTH_STATUS_URI}

echo
echo
echo "Collecting GetHealthStatus metrics ..."
curl --aws-sigv4  "aws:amz:us-east-1:kms-xks-proxy" \
  --user ${ACCESS_KEY_ID}:${SECRET_ACCESS_KEY} \
    -X POST -H "Content-Type:application/json" \
    --data @xks_get_health_status_request.json \
    -w "@curl_metrics.txt" -o encrypt.out -s \
    ${XKS_GET_HEALTH_STATUS_URI}

