package handlers

import (
	"github.com/ryanjarv/yxks/pkg/utils"
	"log"
	"net/http"
	"slices"
)

// HealthHandler endpoint
// URI: /kms/xks/v1/health
func HealthHandler(w http.ResponseWriter, req *http.Request) {
	utils.H[GetHealthStatusRequest](w, req, func(t GetHealthStatusRequest) (any, error) {
		log.Printf("[DEBUG] health handler called with: %s\n", utils.TryJson(t))

		response, err := Health(t)
		if err != nil {
			log.Printf("[ERROR] error encrypting: %v", err)
			return nil, err
		}

		return response, nil
	})
}

func Health(req GetHealthStatusRequest) (any, error) {
	if slices.Contains(KmsOperations, req.RequestMetadata.KmsOperation) {
		log.Printf("[ERROR] unknown KMS operation: %s\n", req.RequestMetadata.KmsOperation)
	}

	return &GetHealthStatusResponse{
		XksProxyFleetSize: 2,
		XksProxyVendor:    "Acme Corp",
		XksProxyModel:     "Acme XKS Proxy 1.0",
		EkmVendor:         "Thales Group",
		EkmFleetDetails: []EkmFleetDetail{
			{
				Id:           "hsm-id-1",
				Model:        "Luna 5.0",
				HealthStatus: HealthStatusACTIVE,
			},
		},
	}, nil
}
