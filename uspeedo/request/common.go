/*
Package request is the request of service
*/
package request

import "time"

// Common is the common request
type Common interface {
	GetAction() string
	SetAction(string) error

	GetActionRef() *string
	SetActionRef(*string) error

	SetRetryCount(int)
	GetRetryCount() int

	WithRetry(int)
	GetMaxretries() int

	WithTimeout(time.Duration)
	GetTimeout() time.Duration

	SetRequestTime(time.Time)
	GetRequestTime() time.Time

	SetRetryable(retryable bool)
	GetRetryable() bool

	SetEncoder(encoder Encoder)
	GetEncoder() Encoder
}

// CommonBase is the base struct of common request
type CommonBase struct {
	Action *string

	maxRetries  int
	retryable   bool
	retryCount  int
	timeout     time.Duration
	requestTime time.Time
	encoder     Encoder
}

// SetRetryCount will set retry count of request
func (c *CommonBase) SetRetryCount(retryCount int) {
	c.retryCount = retryCount
}

// GetRetryCount will return retry count of request
func (c *CommonBase) GetRetryCount() int {
	return c.retryCount
}

// WithRetry will set max retry count of request
func (c *CommonBase) WithRetry(maxRetries int) {
	c.maxRetries = maxRetries
	c.retryable = true
}

// GetMaxretries will return max retry count of request
func (c *CommonBase) GetMaxretries() int {
	return c.maxRetries
}

// SetRetryable will set if the request is retryable
func (c *CommonBase) SetRetryable(retryable bool) {
	c.retryable = retryable
}

// GetRetryable will return if the request is retryable
func (c *CommonBase) GetRetryable() bool {
	return c.retryable
}

// WithTimeout will set timeout of request
func (c *CommonBase) WithTimeout(timeout time.Duration) {
	c.timeout = timeout
}

// GetTimeout will get timeout of request
func (c *CommonBase) GetTimeout() time.Duration {
	return c.timeout
}

// SetRequestTime will set timeout of request
func (c *CommonBase) SetRequestTime(requestTime time.Time) {
	c.requestTime = requestTime
}

// GetRequestTime will get timeout of request
func (c *CommonBase) GetRequestTime() time.Time {
	return c.requestTime
}

// GetAction will return action of request
func (c *CommonBase) GetAction() string {
	if c.Action == nil {
		return ""
	}
	return *c.Action
}

// SetAction will set action of request
func (c *CommonBase) SetAction(val string) error {
	c.Action = &val
	return nil
}

// GetActionRef will return a pointer to action of request
func (c *CommonBase) GetActionRef() *string {
	return c.Action
}

// SetActionRef will set a pointer to action of request
func (c *CommonBase) SetActionRef(val *string) error {
	c.Action = val
	return nil
}

// GetEncoder will return encoder of request
func (c *CommonBase) GetEncoder() Encoder {
	return c.encoder
}

// SetEncoder will set encoder of request
func (c *CommonBase) SetEncoder(encoder Encoder) {
	c.encoder = encoder
}
