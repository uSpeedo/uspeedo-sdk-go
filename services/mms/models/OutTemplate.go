package models

/*
OutTemplate -
*/
type OutTemplate struct {

	/**
	 *
	 */
	TemplateId *string `required:""`
	/**
	 * text content
	 */
	Text *string `required:""`
	/**
	 * properties corresponding to template variables
	 */
	Attributes []VariableAttr `required:""`
	/**
	 *
	 */
	ErrDesc *string `required:""`
	/**
	 * media file address, base64 encoding, 3-second validity period
	 */
	MediaFile *string `required:""`
	/**
	 *
	 */
	Status *int `required:""`
	/**
	 *
	 */
	TemplateName *string `required:""`
	/**
	 *
	 */
	CreateTime *int `required:""`
	/**
	 * media file type
	 */
	MediaType *string `required:""`
	/**
	 * subject
	 */
	Subject *string `required:""`
	/**
	 * tags
	 */
	Tags []int `required:""`
}
