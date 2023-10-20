package http

import (
	"time"
)

const (
	MimeFormURLEncoded = "application/x-www-form-urlencoded"
	MimeJSON           = "application/json"
)

const (
	HeaderNameContentType = "Content-Type"
	HeaderNameUserAgent   = "User-Agent"
	HeaderUTimestampMs    = "U-Timestamp-Ms"

	HeaderNamePublicKey = "X-Access-Key-Id"
	HeaderNameSignature = "X-Signature"
	HeaderNameTimestamp = "X-Timestamp" // 请求时间戳，NTP误差范围5min内
	HeaderNameNonce     = "X-Nonce"     // MAX len 32 char
)

// DefaultHeaders defined default http headers
var DefaultHeaders = map[string]string{
	HeaderNameContentType: MimeJSON,
	// "X-SDK-VERSION": VERSION,
}

// DefaultTimeout is the default timeout of each request
var DefaultTimeout = 30 * time.Second
