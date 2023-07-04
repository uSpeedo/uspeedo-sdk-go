

package models


/*
Console__Component - 
 */
type Console__Component struct {
    
    /**
     * 
     */
    Text *string `required:""`
    /**
     * 
     */
    Type *string `required:""`
    /**
     * 
     */
    Buttons []Console__Button `required:""`
    /**
     * 
     */
    Example *Console__Example `required:""`
    /**
     * 
     */
    Format *string `required:""`
}
