package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the ReplacerConfig construct.
type IReplacerProps interface {
	Replacer() IReplacer
	SetReplacer(r IReplacer)
}

// The jsii proxy for IReplacerProps
type jsiiProxy_IReplacerProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IReplacerProps) Replacer() IReplacer {
	var returns IReplacer
	_jsii_.Get(
		j,
		"replacer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplacerProps)SetReplacer(val IReplacer) {
	if err := j.validateSetReplacerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacer",
		val,
	)
}

