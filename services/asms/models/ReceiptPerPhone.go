package models

/*
ReceiptPerPhone -
*/
type ReceiptPerPhone struct {

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
	UserId *string `required:""`
}
