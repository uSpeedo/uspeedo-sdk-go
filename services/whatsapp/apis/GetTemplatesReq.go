package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
GetTemplatesReq -
*/
type GetTemplatesReq struct {
	request.CommonBase

	/**
	 *
	 */
	After *string `required:""`
	/**
	 *
	 */
	Before *string `required:""`
	/**
	 *
	 */
	BusinessPhone *string `required:"true"`
	/**
	 *
	 */
	Limit *int `required:""`
	/**
	 *
	 */
	Name *string `required:""`
}
