package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
GetUSMSSendReceiptReq -
*/
type GetUSMSSendReceiptReq struct {
	request.CommonBase

	/**
	 *
	 */
	SessionNoSet []string `required:""`
	/**
	 * Account ID,Get accountId reference: https://docs.uspeedo.com/docs/sms/api/
	 */
	AccountId *int `required:""`
}
