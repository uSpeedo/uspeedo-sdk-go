package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
CreateUSMSTemplateReq -
*/
type CreateUSMSTemplateReq struct {
	request.CommonBase

	/**
	 *
	 */
	Remark *string `required:""`
	/**
	 *
	 */
	Template *string `required:"true"`
	/**
	 *
	 */
	TemplateName *string `required:"true"`
	/**
	 * Account ID,Get accountId reference: https://docs.uspeedo.com/docs/sms/api/
	 */
	AccountId *int `required:""`
	/**
	 * Verification code: 1 , Notify : 2, sale : 3
	 */
	Purpose *int `required:"true"`
}
