package asms

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/asms/apis"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
)

// AsmsClient is the client of Asms
type AsmsClient struct {
	*uspeedo.Client
}

// NewClient will return a instance of AsmsClient
func NewClient(config *uspeedo.Config, credential *auth.Credential) *AsmsClient {
	meta := uspeedo.ClientMeta{Product: "Asms"}
	client := uspeedo.NewClientWithMeta(config, credential, meta)
	return &AsmsClient{
		client,
	}
}

// NewQueryUSMSTemplateReq will create request of QueryUSMSTemplate action.
func (c *AsmsClient) NewQueryUSMSTemplateReq() *apis.QueryUSMSTemplateReq {
	req := &apis.QueryUSMSTemplateReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: QueryUSMSTemplate
*/
func (c *AsmsClient) QueryUSMSTemplate(req *apis.QueryUSMSTemplateReq) (*apis.QueryUSMSTemplateResp, error) {
	var err error
	var res apis.QueryUSMSTemplateResp

	err = c.Client.InvokeAction("QueryUSMSTemplate", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewUpdateUSMSTemplateReq will create request of UpdateUSMSTemplate action.
func (c *AsmsClient) NewUpdateUSMSTemplateReq() *apis.UpdateUSMSTemplateReq {
	req := &apis.UpdateUSMSTemplateReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: UpdateUSMSTemplate
*/
func (c *AsmsClient) UpdateUSMSTemplate(req *apis.UpdateUSMSTemplateReq) (*apis.UpdateUSMSTemplateResp, error) {
	var err error
	var res apis.UpdateUSMSTemplateResp

	err = c.Client.InvokeAction("UpdateUSMSTemplate", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewSendBatchUSMSMessageReq will create request of SendBatchUSMSMessage action.
func (c *AsmsClient) NewSendBatchUSMSMessageReq() *apis.SendBatchUSMSMessageReq {
	req := &apis.SendBatchUSMSMessageReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: SendBatchUSMSMessage
*/
func (c *AsmsClient) SendBatchUSMSMessage(req *apis.SendBatchUSMSMessageReq) (*apis.SendBatchUSMSMessageResp, error) {
	var err error
	var res apis.SendBatchUSMSMessageResp

	err = c.Client.InvokeAction("SendBatchUSMSMessage", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewDeleteUSMSTemplateReq will create request of DeleteUSMSTemplate action.
func (c *AsmsClient) NewDeleteUSMSTemplateReq() *apis.DeleteUSMSTemplateReq {
	req := &apis.DeleteUSMSTemplateReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteUSMSTemplate
*/
func (c *AsmsClient) DeleteUSMSTemplate(req *apis.DeleteUSMSTemplateReq) (*apis.DeleteUSMSTemplateResp, error) {
	var err error
	var res apis.DeleteUSMSTemplateResp

	err = c.Client.InvokeAction("DeleteUSMSTemplate", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewCreateUSMSTemplateReq will create request of CreateUSMSTemplate action.
func (c *AsmsClient) NewCreateUSMSTemplateReq() *apis.CreateUSMSTemplateReq {
	req := &apis.CreateUSMSTemplateReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: CreateUSMSTemplate
*/
func (c *AsmsClient) CreateUSMSTemplate(req *apis.CreateUSMSTemplateReq) (*apis.CreateUSMSTemplateResp, error) {
	var err error
	var res apis.CreateUSMSTemplateResp

	err = c.Client.InvokeAction("CreateUSMSTemplate", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewGetUSMSSendReceiptReq will create request of GetUSMSSendReceipt action.
func (c *AsmsClient) NewGetUSMSSendReceiptReq() *apis.GetUSMSSendReceiptReq {
	req := &apis.GetUSMSSendReceiptReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetUSMSSendReceipt
*/
func (c *AsmsClient) GetUSMSSendReceipt(req *apis.GetUSMSSendReceiptReq) (*apis.GetUSMSSendReceiptResp, error) {
	var err error
	var res apis.GetUSMSSendReceiptResp

	err = c.Client.InvokeAction("GetUSMSSendReceipt", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
