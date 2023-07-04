package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
	"unsafe"

	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/config"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/http"
)

type JSONEncoder struct {
	cfg  *config.Config
	cred *auth.Credential
}

// Encode will encode request struct instance as a map for json encoded
func (e *JSONEncoder) Encode(req Common) (*http.HttpRequest, error) {
	if req == nil {
		return nil, fmt.Errorf("invalid request, got nil")
	}

	httpReq := http.NewHttpRequest()
	_ = httpReq.SetURL(e.cfg.BaseUrl)
	_ = httpReq.SetTimeout(req.GetTimeout())
	_ = httpReq.SetMethod("POST")
	_ = httpReq.SetQuery("Action", req.GetAction()) // workaround for http log handler
	_ = httpReq.SetHeader(http.HeaderNameContentType, http.MimeJSON)

	// encode struct to map
	payload, err := EncodeJSON(req)
	if err != nil {
		return nil, err
	}

	// set credential to header
	_ = httpReq.SetHeader(http.HeaderNamePublicKey, e.cred.PublicKey)
	_ = httpReq.SetHeader(http.HeaderNameSignature, e.cred.VerifyAc(payload))
	_ = httpReq.SetHeader(http.HeaderNameTimestamp, strconv.FormatInt(req.GetRequestTime().Unix(), 10))
	_ = httpReq.SetHeader(http.HeaderNameNonce, RandString(10))

	// marshal payload as request body
	bs, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	_ = httpReq.SetRequestBody(bs)
	return httpReq, nil
}

func NewJSONEncoder(cfg *config.Config, cred *auth.Credential) Encoder {
	return &JSONEncoder{cfg: cfg, cred: cred}
}

func EncodeJSON(req Common) (map[string]interface{}, error) {
	return structToMap(req)
}

// encodeOne will encode any value as string
func encodeOne(v *reflect.Value) (string, error) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), nil
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), nil
	case reflect.String:
		return v.String(), nil
	case reflect.Ptr, reflect.Interface:
		ptrValue := v.Elem()
		return encodeOne(&ptrValue)
	default:
		message := fmt.Sprintf(
			"Invalid variable type, type must be one of int-, uint-,"+
				" float-, bool, string and ptr, got %s",
			v.Kind().String(),
		)
		return "", errors.New(message)
	}
}

func RandString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	const (
		letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
