package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/whatsapp/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
UploadMediaResData -
*/
type UploadMediaResData struct {
	response.CommonBase

	/**
	 *
	 */
	Data *models.UploadMediaRes `required:""`
}
