package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
CreateMMSTemplateRes -
*/
type CreateMMSTemplateRes struct {
	response.CommonBase

	/**
	 * template id
	 */
	TemplateId *string `required:""`
}
