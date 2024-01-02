package mms

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/mms/apis"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/auth"
)

// MMSClient is the client of MMS
type MMSClient struct {
	*uspeedo.Client
}

// NewClient will return a instance of MMSClient
func NewClient(config *uspeedo.Config, credential *auth.Credential) *MMSClient {
	meta := uspeedo.ClientMeta{Product: "MMS"}
	client := uspeedo.NewClientWithMeta(config, credential, meta)
	return &MMSClient{
		client,
	}
}

// NewSendBatchMMSMessageReq will create request of SendBatchMMSMessage action.
func (c *MMSClient) NewSendBatchMMSMessageReq() *apis.SendBatchMMSMessageReq {
	req := &apis.SendBatchMMSMessageReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: SendBatchMMSMessage
*/
func (c *MMSClient) SendBatchMMSMessage(req *apis.SendBatchMMSMessageReq) (*apis.SendBatchMMSMessageRes, error) {
	var err error
	var res apis.SendBatchMMSMessageRes

	err = c.Client.InvokeAction("SendBatchMMSMessage", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewQueryMMSTemplateReq will create request of QueryMMSTemplate action.
func (c *MMSClient) NewQueryMMSTemplateReq() *apis.QueryMMSTemplateReq {
	req := &apis.QueryMMSTemplateReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: QueryMMSTemplate
*/
func (c *MMSClient) QueryMMSTemplate(req *apis.QueryMMSTemplateReq) (*apis.QueryMMSTemplateRes, error) {
	var err error
	var res apis.QueryMMSTemplateRes

	err = c.Client.InvokeAction("QueryMMSTemplate", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewDeleteMMSTemplateReq will create request of DeleteMMSTemplate action.
func (c *MMSClient) NewDeleteMMSTemplateReq() *apis.DeleteMMSTemplateReq {
	req := &apis.DeleteMMSTemplateReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteMMSTemplate
*/
func (c *MMSClient) DeleteMMSTemplate(req *apis.DeleteMMSTemplateReq) (*apis.DeleteMMSTemplateRes, error) {
	var err error
	var res apis.DeleteMMSTemplateRes

	err = c.Client.InvokeAction("DeleteMMSTemplate", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewCreateMMSTemplateReq will create request of CreateMMSTemplate action.
func (c *MMSClient) NewCreateMMSTemplateReq() *apis.CreateMMSTemplateReq {
	req := &apis.CreateMMSTemplateReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: CreateMMSTemplate
*/
func (c *MMSClient) CreateMMSTemplate(req *apis.CreateMMSTemplateReq) (*apis.CreateMMSTemplateRes, error) {
	var err error
	var res apis.CreateMMSTemplateRes

	err = c.Client.InvokeAction("CreateMMSTemplate", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewUpdateMMSTemplateReq will create request of UpdateMMSTemplate action.
func (c *MMSClient) NewUpdateMMSTemplateReq() *apis.UpdateMMSTemplateReq {
	req := &apis.UpdateMMSTemplateReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: UpdateMMSTemplate
*/
func (c *MMSClient) UpdateMMSTemplate(req *apis.UpdateMMSTemplateReq) (*apis.UpdateMMSTemplateRes, error) {
	var err error
	var res apis.UpdateMMSTemplateRes

	err = c.Client.InvokeAction("UpdateMMSTemplate", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

// NewGetMMSSendReceiptReq will create request of GetMMSSendReceipt action.
func (c *MMSClient) NewGetMMSSendReceiptReq() *apis.GetMMSSendReceiptReq {
	req := &apis.GetMMSSendReceiptReq{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetMMSSendReceipt
*/
func (c *MMSClient) GetMMSSendReceipt(req *apis.GetMMSSendReceiptReq) (*apis.GetMMSSendReceiptRes, error) {
	var err error
	var res apis.GetMMSSendReceiptRes

	err = c.Client.InvokeAction("GetMMSSendReceipt", req, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
