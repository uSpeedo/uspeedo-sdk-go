package models

/*
FailedTargetEmail -
*/
type FailedTargetEmail struct {

	/**
	 *
	 */
	EmailAddress *string `required:""`
	/**
	 *
	 */
	FailureReason *string `required:""`
	/**
	 * variableName{##}variableValue
	 */
	TemplateVariableParams []string `required:""`
}
