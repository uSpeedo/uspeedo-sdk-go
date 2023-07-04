

package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)


/*
Console__sendMessageResponse - 
 */
type Console__sendMessageResponse struct {
    
    response.CommonBase
    
    /**
     * 
     */
    MessageId *string `required:""`
}
