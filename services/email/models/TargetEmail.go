package models

/*
TargetEmail -
*/
type TargetEmail struct {

	/**
	 *
	 */
	EmailAddress *string `required:""`
	/**
	 * variableName{##}variableValue
	 */
	TemplateVariableParams []string `required:""`
}
