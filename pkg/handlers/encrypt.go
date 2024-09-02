package handlers

import (
	"fmt"
	"github.com/ryanjarv/yxks/pkg/utils"
	"github.com/samber/lo"
	"net/http"
)

// Encrypt endpoint
// URI: /kms/xks/v1/keys/{externalKeyId}/encrypt
func Encrypt(w http.ResponseWriter, req *http.Request) {
	extKeyId := utils.GetExternalKeyId(req)
	lo.Must(fmt.Fprintf(w, "Encrypting with %s: %v", extKeyId, req))
}
