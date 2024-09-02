package handlers

import (
	"fmt"
	"github.com/samber/lo"
	"net/http"
)

// HealthHandler endpoint
// URI: /kms/xks/v1/health
func HealthHandler(w http.ResponseWriter, req *http.Request) {
	resp := `{
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
}`
	lo.Must(fmt.Fprintf(w, "%s", resp))
}
