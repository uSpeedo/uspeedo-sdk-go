

package apis

import (
    "github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)


/*
Console__getMediaRequest - 
 */
type Console__getMediaRequest struct {
    
    request.CommonBase

    
    /**
     * phone number
     */
    BusinessPhone *string `required:"true"`
    /**
     * media id
     */
    MediaId *string `required:"true"`
}
