package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the configuration for waiting during a passive scan.
//
// Example:
//   const passiveScanWaitConfig = new PassiveScanWait({
//     maxDuration: 300
//   });
//
type PassiveScanWait interface {
	IPassiveScanWait
	MaxDuration() *float64
	SetMaxDuration(val *float64)
}

// The jsii proxy struct for PassiveScanWait
type jsiiProxy_PassiveScanWait struct {
	jsiiProxy_IPassiveScanWait
}

func (j *jsiiProxy_PassiveScanWait) MaxDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDuration",
		&returns,
	)
	return returns
}


// Creates an instance of PassiveScanWait.
func NewPassiveScanWait(maxDuration *float64) PassiveScanWait {
	_init_.Initialize()

	j := jsiiProxy_PassiveScanWait{}

	_jsii_.Create(
		"zap-cdk.PassiveScanWait",
		[]interface{}{maxDuration},
		&j,
	)

	return &j
}

// Creates an instance of PassiveScanWait.
func NewPassiveScanWait_Override(p PassiveScanWait, maxDuration *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.PassiveScanWait",
		[]interface{}{maxDuration},
		p,
	)
}

func (j *jsiiProxy_PassiveScanWait)SetMaxDuration(val *float64) {
	_jsii_.Set(
		j,
		"maxDuration",
		val,
	)
}

