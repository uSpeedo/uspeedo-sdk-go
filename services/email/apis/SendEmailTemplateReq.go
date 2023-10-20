package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/services/email/models"
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
SendEmailTemplateReq -
*/
type SendEmailTemplateReq struct {
	request.CommonBase

	/**
	 * account id
	 */
	AccountId *int `required:""`
	/**
	 *
	 */
	EmailContent []models.TargetEmail `required:"true"`
	/**
	 *
	 */
	SendEmail *string `required:"true"`
	/**
	 *
	 */
	TemplateId *string `required:"true"`
}
