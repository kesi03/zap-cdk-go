package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the ImportConfig construct.
type IImportProps interface {
	Import() IImport
	SetImport(i IImport)
}

// The jsii proxy for IImportProps
type jsiiProxy_IImportProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IImportProps) Import() IImport {
	var returns IImport
	_jsii_.Get(
		j,
		"import",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImportProps)SetImport(val IImport) {
	if err := j.validateSetImportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"import",
		val,
	)
}

