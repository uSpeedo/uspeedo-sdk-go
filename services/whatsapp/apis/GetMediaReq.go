package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
GetMediaReq -
*/
type GetMediaReq struct {
	request.CommonBase

	/**
	 * phone number
	 */
	BusinessPhone *string `required:"true"`
	/**
	 * media id
	 */
	MediaId *string `required:"true"`
}
