

package models


/*
Console__TemplateInfo - 
 */
type Console__TemplateInfo struct {
    
    /**
     * 
     */
    Status *string `required:""`
    /**
     * 
     */
    Category *string `required:""`
    /**
     * 
     */
    Language *string `required:""`
    /**
     * 
     */
    RejectedReason *string `required:""`
    /**
     * 
     */
    QualityScore *Console__QualityScore `required:""`
    /**
     * 
     */
    Tags []string `required:""`
    /**
     * 
     */
    Components []Console__Component `required:""`
    /**
     * 
     */
    ID *string `required:""`
    /**
     * 
     */
    Name *string `required:""`
}
