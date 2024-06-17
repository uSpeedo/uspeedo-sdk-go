package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
SendMessageRequest -
*/
type SendMessageRequest struct {
	request.CommonBase

	/**
	 * Account ID,Get accountId reference: https://uspeedo.com/docs/api_sdk/api/
	 */
	AccountId *int `required:"true"`
	/**
	 * Called Phone Number
	 */
	CalledNumber *string `required:"true"`
	/**
	 * Template id
	 */
	TemplateId *string `required:"true"`
}
