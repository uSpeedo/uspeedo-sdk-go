package models

/*
TargetPhoneWithFailure -
*/
type TargetPhoneWithFailure struct {

	/**
	 *
	 */
	TemplateParams []string `required:""`
	/**
	 *
	 */
	FailureDetails *string `required:""`
	/**
	 *
	 */
	Invalid *bool `required:""`
	/**
	 *
	 */
	Phone *string `required:""`
}
