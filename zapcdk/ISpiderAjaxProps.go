package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the SpiderAjaxConfig construct.
type ISpiderAjaxProps interface {
	SpiderAjax() ISpiderAjax
	SetSpiderAjax(s ISpiderAjax)
}

// The jsii proxy for ISpiderAjaxProps
type jsiiProxy_ISpiderAjaxProps struct {
	_ byte // padding
}

func (j *jsiiProxy_ISpiderAjaxProps) SpiderAjax() ISpiderAjax {
	var returns ISpiderAjax
	_jsii_.Get(
		j,
		"spiderAjax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjaxProps)SetSpiderAjax(val ISpiderAjax) {
	if err := j.validateSetSpiderAjaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spiderAjax",
		val,
	)
}

