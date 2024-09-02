package handlers

import (
	"fmt"
	"github.com/ryanjarv/yxks/pkg/utils"
	"github.com/samber/lo"
	"net/http"
)

// Decrypt endpoint
// URI: /kms/xks/v1/keys/{externalKeyId}/decrypt
func Decrypt(w http.ResponseWriter, req *http.Request) {
	extKeyId := utils.GetExternalKeyId(req)
	lo.Must(fmt.Fprintf(w, "Decrypting with %s: %v", extKeyId, req))
}
