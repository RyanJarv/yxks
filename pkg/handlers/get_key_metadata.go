package handlers

import (
	"fmt"
	"github.com/ryanjarv/yxks/pkg/utils"
	"github.com/samber/lo"
	"net/http"
)

// GetKeyMetadata endpoint
// URI: /kms/xks/v1/keys/{externalKeyId}/metadata
func GetKeyMetadata(w http.ResponseWriter, req *http.Request) {
	extKeyId := utils.GetExternalKeyId(req)
	lo.Must(fmt.Fprintf(w, "Getting %s metadata: %+v", extKeyId, req))
}
