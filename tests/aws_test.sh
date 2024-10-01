#!/usr/bin/env bash
# More environment variables in  aws-kms-xksproxy-test-client/utils/test_config.sh

set -euo pipefail

original_plaintext="asdf"

resp=$(
  aws --profile attacker kms encrypt  --key-id arn:aws:kms:us-east-2:123456789012:key/7d2bfabe-63ea-4c05-80cb-89300aa53151 --plaintext "$(echo "$original_plaintext"|base64)"
)
CiphertextBlob=$(
  echo $resp | jq -r '.CiphertextBlob'
)

resp=$(
  aws --profile attacker kms decrypt --key-id arn:aws:kms:us-east-2:123456789012:key/7d2bfabe-63ea-4c05-80cb-89300aa53151 --ciphertext-blob "$CiphertextBlob"
)

decrypted_plaintext=$(
  echo $resp | jq -r '.Plaintext'|base64 -d
)

if [ "$decrypted_plaintext" != "$original_plaintext" ]; then
  echo "Decryption failed"
  exit 1
else
  echo "Decryption succeeded"
fi

