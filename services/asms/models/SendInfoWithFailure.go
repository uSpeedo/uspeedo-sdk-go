package models

/*
SendInfoWithFailure -
*/
type SendInfoWithFailure struct {

	/**
	 * Template ID
	 */
	TemplateId *string `required:""`
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
	Target []TargetPhoneWithFailure `required:""`
}
