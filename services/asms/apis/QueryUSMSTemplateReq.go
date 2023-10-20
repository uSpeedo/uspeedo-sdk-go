package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
QueryUSMSTemplateReq -
*/
type QueryUSMSTemplateReq struct {
	request.CommonBase

	/**
	 * Account ID,Get accountId reference: https://docs.uspeedo.com/docs/sms/api/
	 */
	AccountId *int `required:""`
	/**
	 * Template ID
	 */
	TemplateId []string `required:"true"`
}
