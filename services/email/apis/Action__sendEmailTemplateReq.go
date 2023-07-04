

package apis

import (
    "github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
    "github.com/uspeedo/uspeedo-sdk-go/services/email/models"
)


/*
Action__sendEmailTemplateReq - 
 */
type Action__sendEmailTemplateReq struct {
    
    request.CommonBase

    
    /**
     * account id
     */
    AccountId *int `required:""`
    /**
     * 
     */
    EmailContent []models.Action__TargetEmail `required:"true"`
    /**
     * 
     */
    SendEmail *string `required:"true"`
    /**
     * 
     */
    TemplateId *string `required:"true"`
}
