package whatsapp

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/whatsapp/apis"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
)

// WhatsAppClient is the client of WhatsApp
type WhatsAppClient struct {
	*uspeedo.Client
}

// NewClient will return a instance of WhatsAppClient
func NewClient(config *uspeedo.Config, credential *auth.Credential) *WhatsAppClient {
	meta := uspeedo.ClientMeta{Product: "WhatsApp"}
	client := uspeedo.NewClientWithMeta(config, credential, meta)
	return &WhatsAppClient{
		client,
	}
}

// NewDeleteMediaReq will create request of DeleteMedia action.
func (c *WhatsAppClient) NewDeleteMediaReq() *apis.DeleteMediaReq {
	req := &apis.DeleteMediaReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteMedia
*/
func (c *WhatsAppClient) DeleteMedia(req *apis.DeleteMediaReq) (*apis.DeleteMediaRes, error) {
	var err error
	var res apis.DeleteMediaRes

	err = c.Client.InvokeAction("DeleteMedia", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewUploadMediaReq will create request of UploadMedia action.
func (c *WhatsAppClient) NewUploadMediaReq() *apis.UploadMediaReq {
	req := &apis.UploadMediaReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: UploadMedia
*/
func (c *WhatsAppClient) UploadMedia(req *apis.UploadMediaReq) (*apis.UploadMediaResData, error) {
	var err error
	var res apis.UploadMediaResData

	err = c.Client.InvokeAction("UploadMedia", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewDeleteTemplateReq will create request of DeleteTemplate action.
func (c *WhatsAppClient) NewDeleteTemplateReq() *apis.DeleteTemplateReq {
	req := &apis.DeleteTemplateReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteTemplate
*/
func (c *WhatsAppClient) DeleteTemplate(req *apis.DeleteTemplateReq) (*apis.DeleteTemplateRes, error) {
	var err error
	var res apis.DeleteTemplateRes

	err = c.Client.InvokeAction("DeleteTemplate", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewGetMessageSummaryReq will create request of GetMessageSummary action.
func (c *WhatsAppClient) NewGetMessageSummaryReq() *apis.GetMessageSummaryReq {
	req := &apis.GetMessageSummaryReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetMessageSummary
*/
func (c *WhatsAppClient) GetMessageSummary(req *apis.GetMessageSummaryReq) (*apis.GetMessageSummaryResData, error) {
	var err error
	var res apis.GetMessageSummaryResData

	err = c.Client.InvokeAction("GetMessageSummary", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewGetAccountPhoneListRequest will create request of GetAccountPhoneList action.
func (c *WhatsAppClient) NewGetAccountPhoneListRequest() *apis.GetAccountPhoneListRequest {
	req := &apis.GetAccountPhoneListRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetAccountPhoneList
*/
func (c *WhatsAppClient) GetAccountPhoneList(req *apis.GetAccountPhoneListRequest) (*apis.GetAccountPhoneListResData, error) {
	var err error
	var res apis.GetAccountPhoneListResData

	err = c.Client.InvokeAction("GetAccountPhoneList", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewGetTemplatesReq will create request of GetTemplates action.
func (c *WhatsAppClient) NewGetTemplatesReq() *apis.GetTemplatesReq {
	req := &apis.GetTemplatesReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetTemplates
*/
func (c *WhatsAppClient) GetTemplates(req *apis.GetTemplatesReq) (*apis.GetTemplatesResData, error) {
	var err error
	var res apis.GetTemplatesResData

	err = c.Client.InvokeAction("GetTemplates", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewGetMediaReq will create request of GetMedia action.
func (c *WhatsAppClient) NewGetMediaReq() *apis.GetMediaReq {
	req := &apis.GetMediaReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetMedia
*/
func (c *WhatsAppClient) GetMedia(req *apis.GetMediaReq) (*apis.GetMediaResData, error) {
	var err error
	var res apis.GetMediaResData

	err = c.Client.InvokeAction("GetMedia", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewSendWhatsappMessageReq will create request of SendWhatsappMessage action.
func (c *WhatsAppClient) NewSendWhatsappMessageReq() *apis.SendWhatsappMessageReq {
	req := &apis.SendWhatsappMessageReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: SendWhatsappMessage
*/
func (c *WhatsAppClient) SendWhatsappMessage(req *apis.SendWhatsappMessageReq) (*apis.SendWhatsappMessageResData, error) {
	var err error
	var res apis.SendWhatsappMessageResData

	err = c.Client.InvokeAction("SendWhatsappMessage", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
