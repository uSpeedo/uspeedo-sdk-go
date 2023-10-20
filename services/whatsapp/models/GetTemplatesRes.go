package models

/*
GetTemplatesRes -
*/
type GetTemplatesRes struct {

	/**
	 *
	 */
	Data []TemplateInfo `required:""`
	/**
	 *
	 */
	Paging *GetTemplatesPaging `required:""`
}
