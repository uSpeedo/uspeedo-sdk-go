

package apis

import (
    "github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)


/*
Console__getTemplateRequest - 
 */
type Console__getTemplateRequest struct {
    
    request.CommonBase

    
    /**
     * 
     */
    After *string `required:""`
    /**
     * 
     */
    Before *string `required:""`
    /**
     * 
     */
    BusinessPhone *string `required:"true"`
    /**
     * 
     */
    Limit *int `required:""`
    /**
     * 
     */
    Name *string `required:""`
}
