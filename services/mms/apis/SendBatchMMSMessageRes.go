package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/mms/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
SendBatchMMSMessageRes -
*/
type SendBatchMMSMessageRes struct {
	response.CommonBase

	/**
	 *
	 */
	SessionNo *string `required:""`
	/**
	 *
	 */
	SuccessCount *int `required:""`
	/**
	 *
	 */
	FailContent []models.SendInfoWithFailure `required:""`
}
