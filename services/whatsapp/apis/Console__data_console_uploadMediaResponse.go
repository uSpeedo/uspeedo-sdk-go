

package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
    "github.com/uspeedo/uspeedo-sdk-go/services/whatsapp/models"
)


/*
Console__data_console_uploadMediaResponse - 
 */
type Console__data_console_uploadMediaResponse struct {
    
    response.CommonBase
    
    /**
     * 
     */
    Data *models.Console__uploadMediaResponse `required:""`
}
