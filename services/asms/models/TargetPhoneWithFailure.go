package models

/*
TargetPhoneWithFailure -
*/
type TargetPhoneWithFailure struct {

	/**
	 *
	 */
	Invalid *bool `required:""`
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
	/**
	 *
	 */
	FailureDetails *string `required:""`
}
