package models

/*
ReceiptPerSession -
*/
type ReceiptPerSession struct {

	/**
	 *
	 */
	SessionNo *string `required:""`
	/**
	 *
	 */
	ReceiptSet []ReceiptPerPhone `required:""`
}
