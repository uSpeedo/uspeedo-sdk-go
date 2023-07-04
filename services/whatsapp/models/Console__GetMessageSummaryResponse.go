

package models


/*
Console__GetMessageSummaryResponse - 
 */
type Console__GetMessageSummaryResponse struct {
    
    /**
     * detailed message statistics
     */
    MsgList []Client__MessageSummary `required:""`
    /**
     * total number of messages
     */
    MsgNum *int `required:""`
    /**
     * total amount of messages
     */
    MsgAmount *int `required:""`
}
