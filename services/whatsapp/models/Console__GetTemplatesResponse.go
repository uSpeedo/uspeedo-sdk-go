

package models


/*
Console__GetTemplatesResponse - 
 */
type Console__GetTemplatesResponse struct {
    
    /**
     * 
     */
    Paging *Console__Paging `required:""`
    /**
     * 
     */
    Data []Console__TemplateInfo `required:""`
}
