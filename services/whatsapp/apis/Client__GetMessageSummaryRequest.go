

package apis

import (
    "github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)


/*
Client__GetMessageSummaryRequest - 
 */
type Client__GetMessageSummaryRequest struct {
    
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
