package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
UpdateMMSTemplateReq -
*/
type UpdateMMSTemplateReq struct {
	request.CommonBase

	/**
	 * Account ID,Get accountId reference: https://docs.uspeedo.com/docs/mms/api/
	 */
	AccountId *int `required:"true"`
	/**
	 * media file types
	 */
	MediaType *string `required:""`
	/**
	 * subject
	 */
	Subject *string `required:""`
	/**
	 *
	 */
	TemplateId *string `required:"true"`
	/**
	 *
	 */
	TemplateName *string `required:""`
	/**
	 * media files, base64 encoded
	 */
	Media *string `required:""`
	/**
	 * tags
	 */
	Tags []int `required:""`
	/**
	 *
	 */
	Text *string `required:""`
	/**
	 * properties corresponding to template variables
	 */
	VariableAttr *string `required:""`
}
