package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
DeleteUSMSTemplateReq -
*/
type DeleteUSMSTemplateReq struct {
	request.CommonBase

	/**
	 * Account ID,Get accountId reference: https://docs.uspeedo.com/docs/sms/api/
	 */
	AccountId *int `required:""`
	/**
	 *
	 */
	TemplateIds []string `required:"true"`
}
