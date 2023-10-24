package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/asms/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
SendBatchUSMSMessageResp -
*/
type SendBatchUSMSMessageResp struct {
	response.CommonBase

	/**
	 *
	 */
	SuccessCount *int `required:""`
	/**
	 *
	 */
	FailContent []models.SendInfoWithFailure `required:""`
	/**
	 *
	 */
	SessionNo *string `required:""`
}
