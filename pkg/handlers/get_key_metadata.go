package handlers

import (
	"fmt"
	"github.com/ryanjarv/yxks/pkg/utils"
	"github.com/samber/lo"
	"net/http"
)

// GetKeyMetadataHandler endpoint
// URI: /kms/xks/v1/keys/{externalKeyId}/metadata
func GetKeyMetadataHandler(w http.ResponseWriter, req *http.Request) {
	extKeyId := utils.GetExternalKeyId(req)
	lo.Must(fmt.Fprintf(w, "Getting %s metadata: %+v", extKeyId, req))
}
