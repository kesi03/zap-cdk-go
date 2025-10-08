package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IDataDrivenNode interface {
	Name() *string
	SetName(n *string)
	Regex() *string
	SetRegex(r *string)
}

// The jsii proxy for IDataDrivenNode
type jsiiProxy_IDataDrivenNode struct {
	_ byte // padding
}

func (j *jsiiProxy_IDataDrivenNode) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataDrivenNode)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IDataDrivenNode) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataDrivenNode)SetRegex(val *string) {
	if err := j.validateSetRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regex",
		val,
	)
}

