

package apis

import (
	"github.com/uspeedo/uspeedo-sdk-go/uspeedo/response"
)


/*
Clients__GetMediaResponse - 
 */
type Clients__GetMediaResponse struct {
    
    response.CommonBase
    
    /**
     * 
     */
    file_name *string `required:""`
    /**
     * 
     */
    mime_type *string `required:""`
    /**
     * 
     */
    url *string `required:""`
}
