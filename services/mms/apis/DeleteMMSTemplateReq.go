package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
DeleteMMSTemplateReq -
*/
type DeleteMMSTemplateReq struct {
	request.CommonBase

	/**
	 * Account ID,Get accountId reference: https://docs.uspeedo.com/docs/mms/api/
	 */
	AccountId *int `required:"true"`
	/**
	 *
	 */
	TemplateIds []string `required:"true"`
}
