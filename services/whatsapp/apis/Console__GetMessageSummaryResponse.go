

package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
    "github.com/uspeedo/uspeedo-sdk-go/services/whatsapp/models"
)


/*
Console__GetMessageSummaryResponse - 
 */
type Console__GetMessageSummaryResponse struct {
    
    response.CommonBase
    
    /**
     * total amount of messages
     */
    MsgAmount *int `required:""`
    /**
     * detailed message statistics
     */
    MsgList []models.Client__MessageSummary `required:""`
    /**
     * total number of messages
     */
    MsgNum *int `required:""`
}
