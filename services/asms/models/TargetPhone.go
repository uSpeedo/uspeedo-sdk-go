package models

/*
TargetPhone -
*/
type TargetPhone struct {

	/**
	 *
	 */
	Phone *string `required:""`
	/**
	 *
	 */
	TemplateParams []string `required:""`
	/**
	 *
	 */
	UserId *string `required:""`
	/**
	 *
	 */
	ExtendCode *string `required:""`
}
