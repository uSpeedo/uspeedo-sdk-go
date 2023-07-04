

package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
    "github.com/uspeedo/uspeedo-sdk-go/services/email/models"
)


/*
Action__SendEmailTemplateResponse - 
 */
type Action__SendEmailTemplateResponse struct {
    
    response.CommonBase
    
    /**
     * 
     */
    FailContent []models.Action__FailedTargetEmail `required:""`
    /**
     * The unique identifier of this sending task
     */
    SessionNo *string `required:""`
    /**
     * 
     */
    SuccessCount *int `required:""`
}
