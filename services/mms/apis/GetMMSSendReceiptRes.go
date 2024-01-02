package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/mms/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
GetMMSSendReceiptRes -
*/
type GetMMSSendReceiptRes struct {
	response.CommonBase

	/**
	 *
	 */
	Data []models.ReceiptPerSession `required:""`
}
