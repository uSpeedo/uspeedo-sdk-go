package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/asms/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
GetUSMSSendReceiptResp -
*/
type GetUSMSSendReceiptResp struct {
	response.CommonBase

	/**
	 *
	 */
	Data []models.ReceiptPerSession `required:""`
}
