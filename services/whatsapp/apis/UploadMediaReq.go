package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
UploadMediaReq -
*/
type UploadMediaReq struct {
	request.CommonBase

	/**
	 * phone number
	 */
	BusinessPhone *string `required:"true"`
	/**
	 * base64 encoded raw file, without mime type prefix
	 */
	File *string `required:"true"`
}
