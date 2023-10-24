package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
CreateUSMSTemplateResp -
*/
type CreateUSMSTemplateResp struct {
	response.CommonBase

	/**
	 *
	 */
	TemplateId *string `required:""`
}
