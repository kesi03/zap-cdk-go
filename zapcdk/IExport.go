package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IExport interface {
	Context() *string
	SetContext(c *string)
	FileName() *string
	SetFileName(f *string)
	Source() *string
	SetSource(s *string)
	Type() *string
	SetType(t *string)
}

// The jsii proxy for IExport
type jsiiProxy_IExport struct {
	_ byte // padding
}

func (j *jsiiProxy_IExport) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExport)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_IExport) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExport)SetFileName(val *string) {
	if err := j.validateSetFileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileName",
		val,
	)
}

func (j *jsiiProxy_IExport) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExport)SetSource(val *string) {
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_IExport) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExport)SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

