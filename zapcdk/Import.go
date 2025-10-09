package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

type Import interface {
	IImport
	FileName() *string
	SetFileName(val *string)
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for Import
type jsiiProxy_Import struct {
	jsiiProxy_IImport
}

func (j *jsiiProxy_Import) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Import) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates an instance of Import.
//
// Example:
//   const importConfig = new Import({
//     type: 'har',
//     fileName: 'import.har'
//   });
//
func NewImport(options IImport) Import {
	_init_.Initialize()

	if err := validateNewImportParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Import{}

	_jsii_.Create(
		"zap-cdk.Import",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of Import.
//
// Example:
//   const importConfig = new Import({
//     type: 'har',
//     fileName: 'import.har'
//   });
//
func NewImport_Override(i Import, options IImport) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.Import",
		[]interface{}{options},
		i,
	)
}

func (j *jsiiProxy_Import)SetFileName(val *string) {
	if err := j.validateSetFileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileName",
		val,
	)
}

func (j *jsiiProxy_Import)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

