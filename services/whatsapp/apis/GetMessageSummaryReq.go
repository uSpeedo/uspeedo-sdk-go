package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
GetMessageSummaryReq -
*/
type GetMessageSummaryReq struct {
	request.CommonBase

	/**
	 *
	 */
	AccountId *int `required:""`
	/**
	 *
	 */
	EndTime *int `required:""`
	/**
	 *
	 */
	StartTime *int `required:""`
}
