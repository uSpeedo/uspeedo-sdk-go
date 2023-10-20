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
	FailContent []models.BatchWithFailure `required:""`
	/**
	 *
	 */
	SessionNo *string `required:""`
	/**
	 *
	 */
	SuccessCount *int `required:""`
}
