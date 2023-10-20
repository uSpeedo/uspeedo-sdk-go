package email

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/email/apis"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
)

// EmailClient is the client of Email
type EmailClient struct {
	*uspeedo.Client
}

// NewClient will return a instance of EmailClient
func NewClient(config *uspeedo.Config, credential *auth.Credential) *EmailClient {
	meta := uspeedo.ClientMeta{Product: "Email"}
	client := uspeedo.NewClientWithMeta(config, credential, meta)
	return &EmailClient{
		client,
	}
}

// NewSendEmailTemplateReq will create request of SendEmailTemplate action.
func (c *EmailClient) NewSendEmailTemplateReq() *apis.SendEmailTemplateReq {
	req := &apis.SendEmailTemplateReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: SendEmailTemplate
*/
func (c *EmailClient) SendEmailTemplate(req *apis.SendEmailTemplateReq) (*apis.SendEmailTemplateRes, error) {
	var err error
	var res apis.SendEmailTemplateRes

	err = c.Client.InvokeAction("SendEmailTemplate", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
