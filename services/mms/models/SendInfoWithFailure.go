package models

/*
SendInfoWithFailure -
*/
type SendInfoWithFailure struct {

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
	/**
	 * Template ID
	 */
	TemplateId *string `required:""`
}
