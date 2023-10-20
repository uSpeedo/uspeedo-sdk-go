package models

/*
ReceiptPerPhone -
*/
type ReceiptPerPhone struct {

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
}
