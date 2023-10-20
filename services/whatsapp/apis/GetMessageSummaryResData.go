package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/whatsapp/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
GetMessageSummaryResData -
*/
type GetMessageSummaryResData struct {
	response.CommonBase

	/**
	 *
	 */
	Data *models.GetMessageSummaryRes `required:""`
}
