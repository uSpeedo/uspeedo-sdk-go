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
	TemplateName *string `required:"true"`
	/**
	 * Verification code: 1 , Notify : 2, sale : 3
	 */
	Purpose *int `required:"true"`
	/**
	 *
	 */
	Template *string `required:"true"`
	/**
	 * Account ID,Get accountId reference: https://docs.uspeedo.com/docs/sms/api/
	 */
	AccountId *int `required:""`
	/**
	 *
	 */
	Remark *string `required:""`
}
