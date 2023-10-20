package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/whatsapp/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
GetMediaResData -
*/
type GetMediaResData struct {
	response.CommonBase

	/**
	 *
	 */
	Data *models.GetMediaRes `required:""`
}
