package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type INewType interface {
	Name() *string
	SetName(n *string)
	Parameters() *map[string]interface{}
	SetParameters(p *map[string]interface{})
}

// The jsii proxy for INewType
type jsiiProxy_INewType struct {
	_ byte // padding
}

func (j *jsiiProxy_INewType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INewType)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_INewType) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INewType)SetParameters(val *map[string]interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

