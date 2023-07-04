

package apis

import (
    "github.com/uspeedo/uspeedo-sdk-go/uspeedo/request"
)


/*
Console__uploadMediaRequest - 
 */
type Console__uploadMediaRequest struct {
    
    request.CommonBase

    
    /**
     * phone number
     */
    BusinessPhone *string `required:"true"`
    /**
     * base64 encoded raw file, without mime type prefix
     */
    File *string `required:"true"`
}
