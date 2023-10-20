package models

/*
GetAccountPhoneListResItem -
*/
type GetAccountPhoneListResItem struct {

	/**
	 *
	 */
	Country *string `required:""`
	/**
	 *
	 */
	DisplayName *string `required:""`
	/**
	 *
	 */
	MessagingLimit *string `required:""`
	/**
	 *
	 */
	Number *string `required:""`
	/**
	 *
	 */
	QualityRating *string `required:""`
	/**
	 *
	 */
	Status *string `required:""`
}
