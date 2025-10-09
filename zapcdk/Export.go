package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing an export operation.
//
// Example:
//   const exportConfig = new Export({
//     fileName: 'export.har',
//     type: 'har',
//     source: 'history',
//     context: 'MyContext'
//   });
//
type Export interface {
	IExport
	Context() *string
	SetContext(val *string)
	FileName() *string
	SetFileName(val *string)
	Source() *string
	SetSource(val *string)
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for Export
type jsiiProxy_Export struct {
	jsiiProxy_IExport
}

func (j *jsiiProxy_Export) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Export) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Export) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Export) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates an instance of Export.
func NewExport(options IExport) Export {
	_init_.Initialize()

	if err := validateNewExportParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Export{}

	_jsii_.Create(
		"zap-cdk.Export",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of Export.
func NewExport_Override(e Export, options IExport) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.Export",
		[]interface{}{options},
		e,
	)
}

func (j *jsiiProxy_Export)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_Export)SetFileName(val *string) {
	if err := j.validateSetFileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileName",
		val,
	)
}

func (j *jsiiProxy_Export)SetSource(val *string) {
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_Export)SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

