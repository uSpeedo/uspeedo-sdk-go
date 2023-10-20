package models

/*
GetMediaRes -
*/
type GetMediaRes struct {

	/**
	 *
	 */
	FileName *string `required:""`
	/**
	 *
	 */
	MimeType *string `required:""`
	/**
	 *
	 */
	URL *string `required:""`
}
