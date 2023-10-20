package models

/*
TemplateComponent -
*/
type TemplateComponent struct {

	/**
	 *
	 */
	Example *TemplateComponentExample `required:""`
	/**
	 *
	 */
	Format *string `required:""`
	/**
	 *
	 */
	Text *string `required:""`
	/**
	 *
	 */
	Type *string `required:""`
	/**
	 *
	 */
	Buttons []TemplateComponentButton `required:""`
}
