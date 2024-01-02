package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
CreateMMSTemplateReq -
*/
type CreateMMSTemplateReq struct {
	request.CommonBase

	/**
	 * media files, base64 encoded, less than 300kb
	 */
	Media *string `required:"true"`
	/**
	 * template name
	 */
	TemplateName *string `required:"true"`
	/**
	 * text content
	 */
	Text *string `required:"true"`
	/**
	 * Account ID,Get accountId reference: https://docs.uspeedo.com/docs/mms/api/
	 */
	AccountId *int `required:"true"`
	/**
	 * media file type
	 */
	MediaType *string `required:"true"`
	/**
	 * subject
	 */
	Subject *string `required:""`
}
