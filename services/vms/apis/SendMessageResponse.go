package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)

/*
SendMessageResponse -
*/
type SendMessageResponse struct {
	response.CommonBase

	/**
	 * The unique ID of the send task submitted for this session, which can be used to query the list of SMS messages sent during this session
	 */
	SessionNo *string `required:""`
}
