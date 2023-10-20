package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
ApplyUSMSTemplateResp -
*/
type ApplyUSMSTemplateResp struct {
	response.CommonBase

	/**
	 *
	 */
	TemplateId *string `required:""`
}
