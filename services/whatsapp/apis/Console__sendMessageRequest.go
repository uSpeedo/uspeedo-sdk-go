

package apis

import (
    "github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)


/*
Console__sendMessageRequest - 
 */
type Console__sendMessageRequest struct {
    
    request.CommonBase

    
    /**
     * target phone number
     */
    To *string `required:"true"`
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
}
