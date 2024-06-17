package vms

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/vms/apis"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
)

// VMSClient is the client of VMS
type VMSClient struct {
	*uspeedo.Client
}

// NewClient will return a instance of VMSClient
func NewClient(config *uspeedo.Config, credential *auth.Credential) *VMSClient {
	meta := uspeedo.ClientMeta{Product: "VMS"}
	client := uspeedo.NewClientWithMeta(config, credential, meta)
	return &VMSClient{
		client,
	}
}

// NewSendMessageRequest will create request of SendVMSMessage action.
func (c *VMSClient) NewSendMessageRequest() *apis.SendMessageRequest {
	req := &apis.SendMessageRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: SendVMSMessage
*/
func (c *VMSClient) SendVMSMessage(req *apis.SendMessageRequest) (*apis.SendMessageResponse, error) {
	var err error
	var res apis.SendMessageResponse

	err = c.Client.InvokeAction("SendVMSMessage", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
