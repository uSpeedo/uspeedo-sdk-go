package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
UpdateUSMSTemplateReq -
*/
type UpdateUSMSTemplateReq struct {
	request.CommonBase

	/**
	 *
	 */
	Template *string `required:""`
	/**
	 *
	 */
	TemplateId *string `required:"true"`
	/**
	 *
	 */
	TemplateName *string `required:""`
	/**
	 * Properties corresponding to template variables
	 */
	VariableAttr *string `required:""`
	/**
	 * Account ID,Get accountId reference: https://docs.uspeedo.com/docs/sms/api/
	 */
	AccountId *int `required:""`
	/**
	 *
	 */
	Instruction *string `required:""`
	/**
	 * tags
	 */
	Tags []int `required:""`
}
