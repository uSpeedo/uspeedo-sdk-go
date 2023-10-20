package models

/*
OutTemplate -
*/
type OutTemplate struct {

	/**
	 *
	 */
	Status *int `required:""`
	/**
	 *
	 */
	TemplateId *string `required:""`
	/**
	 * Properties corresponding to template variables
	 */
	Attributes []VariableAttr `required:""`
	/**
	 *
	 */
	CreateTime *int `required:""`
	/**
	 *
	 */
	Remark *string `required:""`
	/**
	 * tags
	 */
	Tags []int `required:""`
	/**
	 *
	 */
	Template *string `required:""`
	/**
	 *
	 */
	TemplateName *string `required:""`
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
}
