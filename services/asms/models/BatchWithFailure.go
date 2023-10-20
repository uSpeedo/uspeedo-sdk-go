package models

/*
BatchWithFailure -
*/
type BatchWithFailure struct {

	/**
	 *
	 */
	FailureDetails *string `required:""`
	/**
	 * Sender Name is basically used to identify the sender of the bulk SMS message.
	 */
	SenderId *string `required:""`
	/**
	 *
	 */
	Target []TargetPhone `required:""`
	/**
	 * Template ID
	 */
	TemplateId *string `required:""`
}
