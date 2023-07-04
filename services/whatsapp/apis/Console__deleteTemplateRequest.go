

package apis

import (
    "github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)


/*
Console__deleteTemplateRequest - 
 */
type Console__deleteTemplateRequest struct {
    
    request.CommonBase

    
    /**
     * 
     */
    BusinessPhone *string `required:"true"`
    /**
     * 
     */
    Name *string `required:"true"`
}
