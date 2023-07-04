package request

import (
	"encoding/base64"

	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/http"
)

type Encoder interface {
	Encode(req Common) (*http.HttpRequest, error)
}

// ToBase64Query will encode a wrapped string as base64 wrapped string
func ToBase64Query(s *string) *string {
	return String(base64.StdEncoding.EncodeToString([]byte(StringValue(s))))
}
