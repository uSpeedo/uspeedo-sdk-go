package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/asms/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
SendBatchUSMSMessageReq -
*/
type SendBatchUSMSMessageReq struct {
	request.CommonBase

	/**
	 * Account ID,Get accountId reference: https://docs.uspeedo.com/docs/sms/api/
	 */
	AccountId *int `required:""`
	/**
	 *
	 */
	TaskContent []models.SendInfo `required:"true"`
}
