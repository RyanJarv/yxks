package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// HealthHandler endpoint
// URI: /kms/xks/v1/health
func HealthHandler(w http.ResponseWriter, req *http.Request) {
	err := healthHandlerErr(w, req)
	if err != nil {
		panic(err)
	}
}

func healthHandlerErr(w http.ResponseWriter, req *http.Request) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("error reading request body: %v", err)
	}

	var encryptReq GetHealthStatusRequest
	if err := json.Unmarshal(body, &encryptReq); err != nil {
		return fmt.Errorf("error unmarshalling request body: %v", err)
	}

	response, err := Health(encryptReq)
	if err != nil {
		return fmt.Errorf("error encrypting: %v", err)
	}

	marshal, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("error marshalling response: %v", err)
	}

	if _, err := w.Write(marshal); err != nil {
		return fmt.Errorf("error writing response: %v", err)
	}

	return nil
}

func Health(req GetHealthStatusRequest) (*GetHealthStatusResponse, error) {
	return &GetHealthStatusResponse{
		XksProxyFleetSize: 2,
		XksProxyVendor:    "Acme Corp",
		XksProxyModel:     "Acme XKS Proxy 1.0",
		EkmVendor:         "Thales Group",
		EkmFleetDetails: []EkmFleetDetail{
			{
				Id:           "hsm-id-1",
				Model:        "Luna 5.0",
				HealthStatus: "DEGRADED",
			},
		},
	}, nil
}
