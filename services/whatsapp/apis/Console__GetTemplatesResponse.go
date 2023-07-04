

package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
    "github.com/uspeedo/uspeedo-sdk-go/services/whatsapp/models"
)


/*
Console__GetTemplatesResponse - 
 */
type Console__GetTemplatesResponse struct {
    
    response.CommonBase
    
    /**
     * 
     */
    Paging *models.Console__Paging `required:""`
    /**
     * 
     */
    Data []models.Console__TemplateInfo `required:""`
}
