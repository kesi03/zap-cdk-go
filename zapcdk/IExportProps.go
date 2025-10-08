package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the ExportConfig construct.
type IExportProps interface {
	Export() IExport
	SetExport(e IExport)
}

// The jsii proxy for IExportProps
type jsiiProxy_IExportProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IExportProps) Export() IExport {
	var returns IExport
	_jsii_.Get(
		j,
		"export",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExportProps)SetExport(val IExport) {
	if err := j.validateSetExportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"export",
		val,
	)
}

