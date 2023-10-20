package models

/*
GetMessageSummaryRes -
*/
type GetMessageSummaryRes struct {

	/**
	 * total amount of messages
	 */
	MsgAmount *int `required:""`
	/**
	 * detailed message statistics
	 */
	MsgList []MessageSummary `required:""`
	/**
	 * total number of messages
	 */
	MsgNum *int `required:""`
}
