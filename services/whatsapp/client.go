
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
// NewConsole__uploadMediaRequest will create request of UploadMedia action.
func (c *WhatsAppClient) NewConsole__uploadMediaRequest() *apis.Console__uploadMediaRequest {
    req := &apis.Console__uploadMediaRequest{}

    // setup request with client config
    c.Client.SetupRequest(req)

    // setup retryable with default retry policy (retry for non-create action and common error)
    req.SetRetryable(true)
    return req
}

/*
API: UploadMedia


*/
func (c *WhatsAppClient) UploadMedia(req *apis.Console__uploadMediaRequest) (*apis.Console__data_console_uploadMediaResponse, error) {
    var err error
    var res apis.Console__data_console_uploadMediaResponse

    err = c.Client.InvokeAction("UploadMedia", req, &res)
    if err != nil {
        return &res, err
    }
    return &res, nil
}
// NewConsole__deleteMediaRequest will create request of DeleteMedia action.
func (c *WhatsAppClient) NewConsole__deleteMediaRequest() *apis.Console__deleteMediaRequest {
    req := &apis.Console__deleteMediaRequest{}

    // setup request with client config
    c.Client.SetupRequest(req)

    // setup retryable with default retry policy (retry for non-create action and common error)
    req.SetRetryable(true)
    return req
}

/*
API: DeleteMedia


*/
func (c *WhatsAppClient) DeleteMedia(req *apis.Console__deleteMediaRequest) (*apis.Common__Empty, error) {
    var err error
    var res apis.Common__Empty

    err = c.Client.InvokeAction("DeleteMedia", req, &res)
    if err != nil {
        return &res, err
    }
    return &res, nil
}
// NewConsole__sendMessageRequest will create request of SendWhatsappMessage action.
func (c *WhatsAppClient) NewConsole__sendMessageRequest() *apis.Console__sendMessageRequest {
    req := &apis.Console__sendMessageRequest{}

    // setup request with client config
    c.Client.SetupRequest(req)

    // setup retryable with default retry policy (retry for non-create action and common error)
    req.SetRetryable(true)
    return req
}

/*
API: SendWhatsappMessage


*/
func (c *WhatsAppClient) SendWhatsappMessage(req *apis.Console__sendMessageRequest) (*apis.Console__data_console_sendMessageResponse, error) {
    var err error
    var res apis.Console__data_console_sendMessageResponse

    err = c.Client.InvokeAction("SendWhatsappMessage", req, &res)
    if err != nil {
        return &res, err
    }
    return &res, nil
}
// NewConsole__getTemplateRequest will create request of GetTemplates action.
func (c *WhatsAppClient) NewConsole__getTemplateRequest() *apis.Console__getTemplateRequest {
    req := &apis.Console__getTemplateRequest{}

    // setup request with client config
    c.Client.SetupRequest(req)

    // setup retryable with default retry policy (retry for non-create action and common error)
    req.SetRetryable(true)
    return req
}

/*
API: GetTemplates


*/
func (c *WhatsAppClient) GetTemplates(req *apis.Console__getTemplateRequest) (*apis.Console__data_console_GetTemplatesResponse, error) {
    var err error
    var res apis.Console__data_console_GetTemplatesResponse

    err = c.Client.InvokeAction("GetTemplates", req, &res)
    if err != nil {
        return &res, err
    }
    return &res, nil
}
// NewConsole__deleteTemplateRequest will create request of DeleteTemplate action.
func (c *WhatsAppClient) NewConsole__deleteTemplateRequest() *apis.Console__deleteTemplateRequest {
    req := &apis.Console__deleteTemplateRequest{}

    // setup request with client config
    c.Client.SetupRequest(req)

    // setup retryable with default retry policy (retry for non-create action and common error)
    req.SetRetryable(true)
    return req
}

/*
API: DeleteTemplate


*/
func (c *WhatsAppClient) DeleteTemplate(req *apis.Console__deleteTemplateRequest) (*apis.Common__Empty, error) {
    var err error
    var res apis.Common__Empty

    err = c.Client.InvokeAction("DeleteTemplate", req, &res)
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
func (c *WhatsAppClient) GetAccountPhoneList(req *apis.GetAccountPhoneListRequest) (*apis.Console__data_console_getAccountAppKeyResponse, error) {
    var err error
    var res apis.Console__data_console_getAccountAppKeyResponse

    err = c.Client.InvokeAction("GetAccountPhoneList", req, &res)
    if err != nil {
        return &res, err
    }
    return &res, nil
}
// NewClient__GetMessageSummaryRequest will create request of GetMessageSummary action.
func (c *WhatsAppClient) NewClient__GetMessageSummaryRequest() *apis.Client__GetMessageSummaryRequest {
    req := &apis.Client__GetMessageSummaryRequest{}

    // setup request with client config
    c.Client.SetupRequest(req)

    // setup retryable with default retry policy (retry for non-create action and common error)
    req.SetRetryable(true)
    return req
}

/*
API: GetMessageSummary


*/
func (c *WhatsAppClient) GetMessageSummary(req *apis.Client__GetMessageSummaryRequest) (*apis.Console__data_console_GetMessageSummaryResponse, error) {
    var err error
    var res apis.Console__data_console_GetMessageSummaryResponse

    err = c.Client.InvokeAction("GetMessageSummary", req, &res)
    if err != nil {
        return &res, err
    }
    return &res, nil
}
// NewConsole__getMediaRequest will create request of GetMedia action.
func (c *WhatsAppClient) NewConsole__getMediaRequest() *apis.Console__getMediaRequest {
    req := &apis.Console__getMediaRequest{}

    // setup request with client config
    c.Client.SetupRequest(req)

    // setup retryable with default retry policy (retry for non-create action and common error)
    req.SetRetryable(true)
    return req
}

/*
API: GetMedia


*/
func (c *WhatsAppClient) GetMedia(req *apis.Console__getMediaRequest) (*apis.Console__data_console_getMediaResponse, error) {
    var err error
    var res apis.Console__data_console_getMediaResponse

    err = c.Client.InvokeAction("GetMedia", req, &res)
    if err != nil {
        return &res, err
    }
    return &res, nil
}
