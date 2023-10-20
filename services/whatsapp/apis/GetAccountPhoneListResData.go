package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/whatsapp/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
GetAccountPhoneListResData -
*/
type GetAccountPhoneListResData struct {
	response.CommonBase

	/**
	 *
	 */
	Data *models.GetAccountPhoneListRes `required:""`
}
