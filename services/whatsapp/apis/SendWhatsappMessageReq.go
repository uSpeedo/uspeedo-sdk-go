package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)

/*
SendWhatsappMessageReq -
*/
type SendWhatsappMessageReq struct {
	request.CommonBase

	/**
	 * message type
	 */
	Type *string `required:"true"`
	/**
	 * send phone number
	 */
	BusinessPhone *string `required:"true"`
	/**
	 * message content in json
	 */
	Content *string `required:"true"`
	/**
	 * target phone number
	 */
	To *string `required:"true"`
}
