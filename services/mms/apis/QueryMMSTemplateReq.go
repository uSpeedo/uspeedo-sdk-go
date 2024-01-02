package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
QueryMMSTemplateReq -
*/
type QueryMMSTemplateReq struct {
	request.CommonBase

	/**
	 * Account ID,Get accountId reference: https://docs.uspeedo.com/docs/mms/api/
	 */
	AccountId *int `required:"true"`
	/**
	 * Template IDs
	 */
	TemplateIds []string `required:"true"`
}
