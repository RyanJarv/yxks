# More environment variables in  aws-kms-xksproxy-test-client/utils/test_config.sh

# Get the local IP address, probably on works on MacOS
local_ip=$(ifconfig|grep 'inet '|grep 'broadcast'|awk '{print $2}')
export XKS_PROXY_HOST=$local_ip

# Change this to the URI_PREFIX of a logical keystore supported by your XKS Proxy.
export URI_PREFIX="example/uri/path/prefix"

# Change this to the Access key ID for request authentication to your logical keystore.
# Valid characters are a-z, A-Z, 0-9, /, - (hyphen), and _ (underscore)
export SIGV4_ACCESS_KEY_ID="BETWEEN2TENAND3TENCHARACTERS"

# Change this to the Secret access key for request authentication to your logical keystore.
# Secret access key must have between 43 and 64 characters. Valid characters are a-z, A-Z, 0-9, /, +, and =
export SIGV4_SECRET_ACCESS_KEY="PleaseReplaceThisWithSomeSecretOfLength43To64"

# Change this to a test key id supported by your logical keystore.
export KEY_ID="foo"

current_dir="$(dirname $(realpath $0))"
cd "$current_dir/aws-kms-xksproxy-test-client"



# Run the tests
./test-xks-proxy-via-docker ENCRYPT_ONLY_KEY_ID
#./test_encrypt_only_key

## Run all the tests including the use of encrypt-only key, decrypt-only key and
## key that can neither encrypt nor decrypt. You can specify the respective key id's with
## the environment variables ENCRYPT_ONLY_KEY_ID, DECRYPT_ONLY_KEY_ID and IMPOTENT_KEY_ID.
#./test-xks-proxy -a
#
## Run all the tests in debug mode, printing the actual curl commands
#DEBUG=1 ./test-xks-proxy
#
## To test against the endpoint http://xks-proxy.mydomain.com
#XKS_PROXY_HOST=xks-proxy.mydomain.com \
#    SCHEME= \
#    ./test-xks-proxy
#
## To enable mTLS, a client side SSL key and certificate would need to be specified.
## The command to run the tests would be something like:
#XKS_PROXY_HOST=xks-proxy_with_mtls_enabled.mydomain.com \
#   MTLS="--key client_key.pem --cert client_cert.pem" \
#   ./test-xks-proxy