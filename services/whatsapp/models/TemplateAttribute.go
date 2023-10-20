package models

/*
TemplateAttribute -
*/
type TemplateAttribute struct {

	/**
	 * named variable map with 1、2
	 */
	NamedVar []string `required:""`
	/**
	 * component item type
	 */
	Type *string `required:""`
}
