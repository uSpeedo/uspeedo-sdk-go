

package models


/*
Action__FailedTargetEmail - 
 */
type Action__FailedTargetEmail struct {
    
    /**
     * variableName{##}variableValue
     */
    TemplateVariableParams []string `required:""`
    /**
     * 
     */
    EmailAddress *string `required:""`
    /**
     * 
     */
    FailureReason *string `required:""`
}
