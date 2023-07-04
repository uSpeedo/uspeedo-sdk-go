

package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
    "github.com/uspeedo/uspeedo-sdk-go/services/whatsapp/models"
)


/*
Console__data_console_getMediaResponse - 
 */
type Console__data_console_getMediaResponse struct {
    
    response.CommonBase
    
    /**
     * 
     */
    Data *models.Console__getMediaResponse `required:""`
}
