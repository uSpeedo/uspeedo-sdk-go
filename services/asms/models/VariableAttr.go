package models

/*
VariableAttr -
*/
type VariableAttr struct {

	/**
	 * 模版参数对应的属性，如果没有属性，则为空
	 */
	Attribute *string `required:""`
	/**
	 * 属性不存在的时候的默认值
	 */
	Default *string `required:""`
	/**
	 * 模版参数名称
	 */
	Variable *string `required:""`
}
