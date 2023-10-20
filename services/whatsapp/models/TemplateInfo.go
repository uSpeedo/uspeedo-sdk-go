package models

/*
TemplateInfo -
*/
type TemplateInfo struct {

	/**
	 *
	 */
	Category *string `required:""`
	/**
	 *
	 */
	ID *string `required:""`
	/**
	 *
	 */
	QualityScore *TemplateQualityScore `required:""`
	/**
	 *
	 */
	Status *string `required:""`
	/**
	 *
	 */
	Tags []string `required:""`
	/**
	 *
	 */
	Attributes []TemplateAttribute `required:""`
	/**
	 *
	 */
	Components []TemplateComponent `required:""`
	/**
	 *
	 */
	Language *string `required:""`
	/**
	 *
	 */
	Name *string `required:""`
	/**
	 *
	 */
	RejectedReason *string `required:""`
}
