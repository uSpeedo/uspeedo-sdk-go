package models

/*
ReceiptPerSession -
*/
type ReceiptPerSession struct {

	/**
	 *
	 */
	ReceiptSet []ReceiptPerPhone `required:""`
	/**
	 *
	 */
	SessionNo *string `required:""`
}
