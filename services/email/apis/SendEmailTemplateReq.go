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
	/**
	 * Subject - 邮件主题（可选）
	 */
	Subject *string `required:""`
	/**
	 * Abstract - 邮件摘要（可选）
	 */
	Abstract *string `required:""`
}
