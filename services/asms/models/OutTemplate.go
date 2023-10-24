package models

/*
OutTemplate -
*/
type OutTemplate struct {

	/**
	 *
	 */
	CreateTime *int `required:""`
	/**
	 *
	 */
	Template *string `required:""`
	/**
	 * Properties corresponding to template variables
	 */
	Attributes []VariableAttr `required:""`
	/**
	 *
	 */
	ErrDesc *string `required:""`
	/**
	 *
	 */
	Instruction *string `required:""`
	/**
	 *
	 */
	Purpose *int `required:""`
	/**
	 *
	 */
	Remark *string `required:""`
	/**
	 *
	 */
	Status *int `required:""`
	/**
	 * tags
	 */
	Tags []int `required:""`
	/**
	 *
	 */
	TemplateId *string `required:""`
	/**
	 *
	 */
	TemplateName *string `required:""`
}
