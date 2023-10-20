package models

/*
GetTemplatesPaging -
*/
type GetTemplatesPaging struct {

	/**
	 *
	 */
	Cursors *GetTemplatesCursors `required:""`
	/**
	 *
	 */
	Next *string `required:""`
	/**
	 *
	 */
	Previous *string `required:""`
}
