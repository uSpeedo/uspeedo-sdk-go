

package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
    "github.com/uspeedo/uspeedo-sdk-go/services/whatsapp/models"
)


/*
Console__data_console_sendMessageResponse - 
 */
type Console__data_console_sendMessageResponse struct {
    
    response.CommonBase
    
    /**
     * 
     */
    Data *models.Console__sendMessageResponse `required:""`
}
