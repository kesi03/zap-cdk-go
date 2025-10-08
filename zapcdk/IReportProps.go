package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the ReportConfig construct.
type IReportProps interface {
	Report() IReport
	SetReport(r IReport)
}

// The jsii proxy for IReportProps
type jsiiProxy_IReportProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IReportProps) Report() IReport {
	var returns IReport
	_jsii_.Get(
		j,
		"report",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportProps)SetReport(val IReport) {
	if err := j.validateSetReportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"report",
		val,
	)
}

