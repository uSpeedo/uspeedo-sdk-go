package models

/*
ReceiptPerPhone -
*/
type ReceiptPerPhone struct {

	/**
	 *
	 */
	ReceiptDesc *string `required:""`
	/**
	 *
	 */
	ReceiptResult *string `required:""`
	/**
	 *
	 */
	ReceiptTime *int `required:""`
	/**
	 *
	 */
	AccountId *int `required:""`
	/**
	 *
	 */
	CostCount *int `required:""`
	/**
	 *
	 */
	Phone *string `required:""`
	/**
	 *
	 */
	ReceiptCode *string `required:""`
}
