package models

/*
VariableAttr -
*/
type VariableAttr struct {

	/**
	 * default value when attribute does not exist
	 */
	Default *string `required:""`
	/**
	 * template parameter name
	 */
	Variable *string `required:""`
	/**
	 * the attribute corresponding to the template parameter, if there is no attribute, it is empty
	 */
	Attribute *string `required:""`
}
