package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IImport interface {
	FileName() *string
	SetFileName(f *string)
	Type() *string
	SetType(t *string)
}

// The jsii proxy for IImport
type jsiiProxy_IImport struct {
	_ byte // padding
}

func (j *jsiiProxy_IImport) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImport)SetFileName(val *string) {
	if err := j.validateSetFileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileName",
		val,
	)
}

func (j *jsiiProxy_IImport) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImport)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

