# More environment variables in  aws-kms-xksproxy-test-client/utils/test_config.sh

# Get the local IP address, probably on works on MacOS
local_ip=$(ifconfig|grep 'inet '|grep 'broadcast'|awk '{print $2}')
export XKS_PROXY_HOST="${local_ip}:8080"
export SCHEME="http://"

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
./test-xks-proxy-via-docker