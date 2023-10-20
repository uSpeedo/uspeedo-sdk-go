package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/email/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
SendEmailTemplateRes -
*/
type SendEmailTemplateRes struct {
	response.CommonBase

	/**
	 * The unique identifier of this sending task
	 */
	SessionNo *string `required:""`
	/**
	 *
	 */
	SuccessCount *int `required:""`
	/**
	 *
	 */
	FailContent []models.FailedTargetEmail `required:""`
}
