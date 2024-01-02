package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/mms/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
SendBatchMMSMessageReq -
*/
type SendBatchMMSMessageReq struct {
	request.CommonBase

	/**
	 * Account ID,Get accountId reference: https://docs.uspeedo.com/docs/mms/api/
	 */
	AccountId *int `required:"true"`
	/**
	 *
	 */
	TaskContent []models.SendInfo `required:"true"`
}
