package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IActiveScanJob interface {
	ActiveScan() IActiveScan
	SetActiveScan(a IActiveScan)
}

// The jsii proxy for IActiveScanJob
type jsiiProxy_IActiveScanJob struct {
	_ byte // padding
}

func (j *jsiiProxy_IActiveScanJob) ActiveScan() IActiveScan {
	var returns IActiveScan
	_jsii_.Get(
		j,
		"activeScan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanJob)SetActiveScan(val IActiveScan) {
	if err := j.validateSetActiveScanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeScan",
		val,
	)
}

