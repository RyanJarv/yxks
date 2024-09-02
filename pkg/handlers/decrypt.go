package handlers

import (
	"fmt"
	"github.com/ryanjarv/yxks/pkg/utils"
	"github.com/samber/lo"
	"net/http"
)

// DecryptHandler endpoint
// URI: /kms/xks/v1/keys/{externalKeyId}/decrypt
func DecryptHandler(w http.ResponseWriter, req *http.Request) {
	extKeyId := utils.GetExternalKeyId(req)
	lo.Must(fmt.Fprintf(w, "Decrypting with %s: %v", extKeyId, req))
}
