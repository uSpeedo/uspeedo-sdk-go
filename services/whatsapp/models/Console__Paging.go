

package models


/*
Console__Paging - 
 */
type Console__Paging struct {
    
    /**
     * 
     */
    cursors *Console__Cursors `required:""`
    /**
     * 
     */
    next *string `required:""`
    /**
     * 
     */
    previous *string `required:""`
}
