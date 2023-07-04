

package models


/*
Action__TargetEmail - 
 */
type Action__TargetEmail struct {
    
    /**
     * 
     */
    EmailAddress *string `required:""`
    /**
     * variableName{##}variableValue
     */
    TemplateVariableParams []string `required:""`
}
