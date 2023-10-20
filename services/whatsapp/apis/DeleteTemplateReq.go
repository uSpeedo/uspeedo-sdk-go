package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
DeleteTemplateReq -
*/
type DeleteTemplateReq struct {
	request.CommonBase

	/**
	 *
	 */
	BusinessPhone *string `required:"true"`
	/**
	 *
	 */
	Name *string `required:"true"`
}
