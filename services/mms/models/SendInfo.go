package models

/*
SendInfo -
*/
type SendInfo struct {

	/**
	 * Sender Name is basically used to identify the sender of the bulk MMS message.
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
