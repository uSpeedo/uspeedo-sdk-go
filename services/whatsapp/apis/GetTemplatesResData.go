package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/whatsapp/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
GetTemplatesResData -
*/
type GetTemplatesResData struct {
	response.CommonBase

	/**
	 *
	 */
	Data *models.GetTemplatesRes `required:""`
}
